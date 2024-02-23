package main

import (
	"github.com/OrigamiWang/msd/micro/framework"
	"github.com/OrigamiWang/msd/register-center/biz"
)

func main() {
	biz.ListenServiceDiscovery()
	biz.ListenHeartBeat()
	root := framework.NewGinWeb()
	// r := root.Group("/")
	{
		// r.GET("/register", mw.PostHandler(handler.GetSvcRegisterListHandler))
		// r.GET("/register/:name", mw.PostHandler(handler.GetSvcRegisterHandler))
	}
	root.Run(":8000")
}
