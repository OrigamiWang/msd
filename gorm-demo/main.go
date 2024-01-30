package main

import (
	"fmt"

	"github.com/OrigamiWang/msd/gorm-demo/handler"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
)

func main() {
	fmt.Println("Hello, world.")
	root := framework.NewGinWeb()
	r := root.Group("/")
	// pprof 性能监
	//pprof.Register(root.Engine)

	d := root.Group("/debug")
	{
		d.GET("/test", mw.PostHandler(handler.Test))

	}
	{
		r.GET("/user", mw.PostHandler(handler.GetUserListHandler))
		r.GET("/user/:id", mw.PostHandler(handler.GetUserByIdHandler))
		r.PUT("/user/:id", mw.PostHandler(handler.UpdateUserHandler, handler.UserBinder))
		r.POST("/user", mw.PostHandler(handler.AddUserHandler, handler.UserBinder))
	}
	root.Run("localhost:8081")
}
