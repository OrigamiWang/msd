package main

import (
	"net/http"

	"github.com/OrigamiWang/msd/manage/biz/kafka"
	"github.com/OrigamiWang/msd/manage/handler"
	"github.com/OrigamiWang/msd/micro/auth/tls"
	"github.com/OrigamiWang/msd/micro/const/svc"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

func init() {
}
func main() {

	kafka.InitKafkaConsumer(svc.MANAGE)
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
	tlsConfig := tls.TlsServerConfig
	srv := &http.Server{
		Addr:      ":8081",
		Handler:   root,
		TLSConfig: tlsConfig,
	}
	if err := srv.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
		logutil.Error("start server failed, err: %v", err)
	}

}
