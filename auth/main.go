package main

import (
	"github.com/OrigamiWang/msd/gate/handler"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
)

// go-jwt

func main() {
	root := framework.NewGinWeb()
	r := root.Group("/")
	d := root.Group("/debug")
	{
		d.GET("/test", mw.PostHandler(handler.TestHandler))
	}
	{
		r.POST("/authorize", mw.PostHandler(handler.AuthorizeHandler))
	}

	root.Run("0.0.0.0:8082")
}
