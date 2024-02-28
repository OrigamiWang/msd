package biz

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/OrigamiWang/msd/gate/biz/lb"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
)

var (
	c_temp           *gin.Context
	sb_temp          *lb.ServiceBalancer
	lb_temp          *lb.LoadBalancer
	client_temp      *http.Client
	serviceName_temp string
)

func Proxy(lb *lb.LoadBalancer, prefix string, group *gin.RouterGroup, client *http.Client) {
	serviceName := strings.TrimPrefix(prefix, "/")
	proxyHandler := func(c *gin.Context) {
		sb := lb.GetServiceBalancer(serviceName)
		if sb == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
			return
		}
		c_temp = c
		sb_temp = sb
		lb_temp = lb
		client_temp = client
		serviceName_temp = serviceName

		instance := sb.GetNextInstance()
		if instance == nil {
			// code: 503, msg: no instance available
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "no instance available"})
			return
		}
		svcUrl, _ := url.Parse(instance.URL)
		proxy := httputil.NewSingleHostReverseProxy(svcUrl)
		proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
			logutil.Error("req: %v, proxy error: %v", req.URL, err)
			currentInstance := sb.GetCurrentInstance()
			logutil.Info("currentInstance: %v", currentInstance)
			if currentInstance == nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{"error": "no instance available"})
				return
			}
			svcUrlCurrent, _ := url.Parse(currentInstance.URL)
			u := fmt.Sprintf("https://%s", svcUrlCurrent.Host)

			// remove offline service
			lb.RemoveOfflineService(serviceName, u)
			// retry another
			nextInstance := sb.GetNextInstance()
			logutil.Info("nextInstance: %v", nextInstance)
			if nextInstance == nil {
				// code: 503, msg: no instance available
				c.JSON(http.StatusServiceUnavailable, gin.H{"error": "no instance available"})
				return
			}
			svcUrlNext, _ := url.Parse(nextInstance.URL)
			proxy := httputil.NewSingleHostReverseProxy(svcUrlNext)
			proxy.ErrorHandler = errHander
			proxy.Transport = client.Transport
			c.Request.URL.Path = c.Param("proxyPath")
			proxy.ServeHTTP(c.Writer, c.Request)
			rw.WriteHeader(http.StatusOK)
		}

		proxy.Transport = client.Transport

		c.Request.URL.Path = c.Param("proxyPath")
		proxy.ServeHTTP(c.Writer, c.Request)
	}
	// avoid repeated proxy registration
	hasProxy := lb.GetHasProxy(serviceName)
	logutil.Info("serviceName: %v, has proxy: %v", serviceName, hasProxy)
	if hasProxy {
		return
	}
	lb.SetHasProxy(serviceName, true)

	// proxy registration
	proxyPath := fmt.Sprintf("%s/*proxyPath", prefix)
	group.GET(proxyPath, proxyHandler)
	group.PUT(proxyPath, proxyHandler)
	group.POST(proxyPath, proxyHandler)
	group.DELETE(proxyPath, proxyHandler)
}

func errHander(rw http.ResponseWriter, req *http.Request, err error) {
	logutil.Error("req: %v, proxy error: %v", req.URL, err)
	currentInstance := sb_temp.GetCurrentInstance()
	logutil.Info("currentInstance: %v", currentInstance)
	if currentInstance == nil {
		c_temp.JSON(http.StatusServiceUnavailable, gin.H{"error": "no instance available"})
		return
	}
	svcUrlCurrent, _ := url.Parse(currentInstance.URL)
	u := fmt.Sprintf("https://%s", svcUrlCurrent.Host)

	// remove offline service
	lb_temp.RemoveOfflineService(serviceName_temp, u)
	// retry another
	nextInstance := sb_temp.GetNextInstance()
	logutil.Info("nextInstance: %v", nextInstance)
	if nextInstance == nil {
		// code: 503, msg: no instance available
		c_temp.JSON(http.StatusServiceUnavailable, gin.H{"error": "no instance available"})
		return
	}
	svcUrlNext, _ := url.Parse(nextInstance.URL)
	proxy := httputil.NewSingleHostReverseProxy(svcUrlNext)
	proxy.ErrorHandler = errHander
	proxy.Transport = client_temp.Transport
	c_temp.Request.URL.Path = c_temp.Param("proxyPath")
	proxy.ServeHTTP(c_temp.Writer, c_temp.Request)
	rw.WriteHeader(http.StatusOK)
}
