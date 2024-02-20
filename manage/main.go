package main

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"

	"github.com/OrigamiWang/msd/manage/biz"
	"github.com/OrigamiWang/msd/manage/biz/kafka"
	"github.com/OrigamiWang/msd/manage/handler"
	"github.com/OrigamiWang/msd/micro/const/svc"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

func init() {
	// init conf
	biz.Init(svc.MANAGE)
}
func main() {
	caCert, _ := os.ReadFile("conf/ca.crt")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	manageCert, _ := tls.LoadX509KeyPair("conf/manage.crt", "conf/manage.key")
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{manageCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
	}

	kafka.InitKafkaConsumer(svc.MANAGE)
	root := framework.NewGinWeb()
	r := root.Group("/")
	// pprof 性能监
	//pprof.Register(root.Engine)

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

	srv := &http.Server{
		Addr:      ":8081",
		Handler:   root,
		TLSConfig: tlsConfig,
	}
	if err := srv.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
		logutil.Error("start server failed, err: %v", err)
	}

}

// func main() {
// 	kafka.InitKafkaConsumer(svc.MANAGE)
// 	root := framework.NewGinWeb()
// 	r := root.Group("/")
// 	// pprof 性能监
// 	//pprof.Register(root.Engine)

// 	d := root.Group("/debug")
// 	{
// 		d.GET("/test", mw.PostHandler(handler.Test))
// 		d.GET("/ext/:key", mw.PostHandler(handler.GetConfExtHandler)) // 用于测试conf-center动态更新配置
// 	}
// 	{
// 		r.GET("/user", mw.PostHandlerWithJwt(handler.GetUserListHandler))
// 		r.GET("/user/:id", mw.PostHandlerWithJwt(handler.GetUserByIdHandler))
// 		r.PUT("/user/:id", mw.PostHandlerWithJwt(handler.UpdateUserHandler, handler.UserBinder))
// 		// login and regiser is no need of jwt
// 		r.POST("/register", mw.PostHandler(handler.AddUserHandler, handler.UserBinder))
// 		r.POST("/login", mw.PostHandler(handler.LoginHandler, handler.LoginBinder))

// 	}
// 	root.Run("0.0.0.0:8081")
// }
