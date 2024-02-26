package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/OrigamiWang/msd/manage/biz/kafka"
	"github.com/OrigamiWang/msd/manage/handler"
	"github.com/OrigamiWang/msd/micro/auth/tls"
	"github.com/OrigamiWang/msd/micro/const/svc"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/OrigamiWang/msd/register-center/biz"
)

func init() {
}

func main() {
	kafka.InitConfCenterConsumer()
	root := framework.NewGinWeb()
	r := root.Group("/")
	// pprof 性能监控
	// pprof.Register(root.Engine)

	d := root.Group("/debug")
	{
		d.GET("/test", mw.PostHandler(handler.Test))
		d.GET("/ext/:key", mw.PostHandler(handler.GetConfExtHandler)) // 用于测试conf-center动态更新配置
	}
	{
		r.GET("/user", mw.PostHandlerWithJwt(handler.GetUserListHandler))
		r.GET("/user/:id", mw.PostHandlerWithJwt(handler.GetUserByIdHandler))
		r.PUT("/user/:id", mw.PostHandlerWithJwt(handler.UpdateUserHandler, handler.UserBinder))
		// login and regiser is no need of jwt
		r.POST("/register", mw.PostHandler(handler.AddUserHandler, handler.UserBinder))
		r.POST("/login", mw.PostHandler(handler.LoginHandler, handler.LoginBinder))

	}

	// get ip, port, instance_id by shell environment variable
	var ip string
	var port int
	var instanceId int
	flag.StringVar(&ip, "ip", "127.0.0.1", "ip")
	flag.IntVar(&port, "port", 8081, "port")
	flag.IntVar(&instanceId, "instance_id", 0, "instance id")
	flag.Parse()
	hostMap := make(map[string]interface{})
	hostMap["ip"] = ip
	hostMap["port"] = port
	hostMap["instance_id"] = instanceId
	jsonBytes, _ := json.Marshal(&hostMap)
	biz.BeatHeartBeat(svc.MANAGE, string(jsonBytes))
	addr := fmt.Sprintf(":%v", port)

	// get tls config and run server
	tlsConfig := tls.TlsServerConfig
	srv := &http.Server{
		Addr:      addr,
		Handler:   root,
		TLSConfig: tlsConfig,
	}
	if err := srv.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
		logutil.Error("start server failed, err: %v", err)
	}

}
