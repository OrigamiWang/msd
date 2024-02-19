package main

import (
	"github.com/OrigamiWang/msd/manage/biz"
	"github.com/OrigamiWang/msd/manage/biz/kafka"
	"github.com/OrigamiWang/msd/manage/handler"
	"github.com/OrigamiWang/msd/micro/const/svc"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
)

func init() {
	// init conf
	biz.Init(svc.MANAGE)
}
func main() {
	kafka.InitKafkaConsumer()
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
	root.Run("0.0.0.0:8081")
}
