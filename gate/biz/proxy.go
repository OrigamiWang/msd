package biz

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/OrigamiWang/msd/micro/const/db"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/OrigamiWang/msd/register-center/model/dao"
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

func GetLiveSvc() {
	var svc_cnt int64 = 10
	match := db.HEART_BEAT_REDIS_PREFIX + "*"
	if keys, _, err := dao.RC.Scan(match, svc_cnt); err == nil {
		for _, key := range keys {
			if val, err := dao.RC.Get(key); err == nil {
				logutil.Info("key: %v, val: %v", key, val)
			} else {
				logutil.Error("get redis error: %v, err")
			}
		}
	} else {
		logutil.Error("scan redis error: %v, err")
	}
}
