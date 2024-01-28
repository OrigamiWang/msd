package main

import (
	"fmt"
	"github.com/OrigamiWang/msd/conf-center/handler"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
)

func main() {
	fmt.Println("conf-center")
	root := framework.NewGinWeb()
	//r := root.Group("/")
	d := root.Group("/debug")
	{
		d.GET("/test", mw.PostHandler(handler.TestHandler))
	}
	{
		//r.GET("/config/:name", mw.PostHandler())
	}
	root.Run()
}
