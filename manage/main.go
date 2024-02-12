package main

import (
	"github.com/OrigamiWang/msd/manage/handler"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
)

func main() {
	root := framework.NewGinWeb()
	r := root.Group("/")
	// pprof 性能监
	//pprof.Register(root.Engine)

	d := root.Group("/debug")
	{
		d.GET("/test", mw.PostHandler(handler.Test))

	}
	{
		r.GET("/user", mw.PostHandlerWithJwt(handler.GetUserListHandler))
		r.GET("/user/:id", mw.PostHandlerWithJwt(handler.GetUserByIdHandler))
		r.PUT("/user/:id", mw.PostHandlerWithJwt(handler.UpdateUserHandler, handler.UserBinder))
		// login and regiser is no need of jwt
		r.POST("/register", mw.PostHandler(handler.AddUserHandler, handler.UserBinder))
		r.POST("/login", mw.PostHandler(handler.LoginHandler, handler.LoginBinder))

	}
	root.Run("localhost:8081")
}
