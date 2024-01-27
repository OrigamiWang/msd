package main

import (
	"fmt"
	"github.com/OrigamiWang/msd/gorm-demo/handler"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
	"github.com/gin-contrib/pprof"
)

func main() {
	fmt.Println("Hello, world.")
	root := framework.NewGinWeb()
	r := root.Group("/")
	pprof.Register(root.Engine)

	d := root.Group("/debug")
	{
		d.GET("/test", mw.PostHandler(handler.Test))

	}
	{
		r.POST("/p", mw.PostHandler(handler.GetFirstUser, handler.UserBinder))
	}
	root.Run()
}
