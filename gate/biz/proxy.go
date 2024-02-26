package biz

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/OrigamiWang/msd/gate/biz/lb"
	"github.com/gin-gonic/gin"
)

func Proxy(lb *lb.LoadBalancer, prefix string, group *gin.RouterGroup, client *http.Client) {
	proxyHandler := func(c *gin.Context) {
		serviceName := strings.TrimPrefix(prefix, "/")
		sb := lb.GetServiceBalancer(serviceName)
		if sb == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
			return
		}

		instance := sb.GetNextInstance()
		svcUrl, _ := url.Parse(instance.URL)
		proxy := httputil.NewSingleHostReverseProxy(svcUrl)
		proxy.Transport = client.Transport

		c.Request.URL.Path = c.Param("proxyPath")
		proxy.ServeHTTP(c.Writer, c.Request)
	}

	proxyPath := fmt.Sprintf("%s/*proxyPath", prefix)
	group.GET(proxyPath, proxyHandler)
	group.PUT(proxyPath, proxyHandler)
	group.POST(proxyPath, proxyHandler)
	group.DELETE(proxyPath, proxyHandler)
}
