package main

import (
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
	"github.com/OrigamiWang/msd/register-center/handler"
)

func main() {
	root := framework.NewGinWeb()
	r := root.Group("/")
	{
		r.GET("/register", mw.PostHandler(handler.GetSvcRegisterListHandler))
		r.GET("/register/:name", mw.PostHandler(handler.GetSvcRegisterHandler))
	}
	root.Run(":8000")
}
