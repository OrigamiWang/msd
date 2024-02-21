package biz

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Proxy(u, prefix string, group *gin.RouterGroup, client *http.Client) {
	svcUrl, _ := url.Parse(u)
	proxy := httputil.NewSingleHostReverseProxy(svcUrl)
	proxy.Transport = client.Transport
	{
		proxyPath := fmt.Sprintf("%s/*proxyPath", prefix)
		proxyHandler := func(c *gin.Context) {
			c.Request.URL.Path = c.Param("proxyPath")
			proxy.ServeHTTP(c.Writer, c.Request)
		}
		group.GET(proxyPath, proxyHandler)
		group.PUT(proxyPath, proxyHandler)
		group.POST(proxyPath, proxyHandler)
		group.DELETE(proxyPath, proxyHandler)
	}
}
