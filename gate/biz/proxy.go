package biz

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Proxy(u, prefix string, group *gin.RouterGroup) {
	svcUrl, _ := url.Parse(u)
	proxy := httputil.NewSingleHostReverseProxy(svcUrl)
	{
		proxyPath := fmt.Sprintf("%s/*proxyPath", prefix)
		group.Any(proxyPath, func(c *gin.Context) {
			c.Request.URL.Path = c.Param("proxyPath")
			proxy.ServeHTTP(c.Writer, c.Request)
		})
	}
}
