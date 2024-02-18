package main

import (
	"fmt"

	"github.com/OrigamiWang/msd/conf-center/dal"
	"github.com/OrigamiWang/msd/conf-center/handler"
	"github.com/OrigamiWang/msd/micro/confparser"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
	"github.com/OrigamiWang/msd/micro/model/dao"
)

func init() {
	// conf-center get config by conf.yml
	// other svc get conf by conf-center
	confparser.Conf = confparser.LoadConf()
	dao.InitDb()
	dal.InitConn()
}

func main() {
	fmt.Println("conf-center")
	root := framework.NewGinWeb()
	r := root.Group("/")
	d := root.Group("/debug")
	{
		d.GET("/test", mw.PostHandler(handler.TestHandler))
	}
	{
		r.GET("/config/:name", handler.GetConfigHandler)
		r.PUT("/config/:name", mw.PostHandler(handler.UpdateConfigHandler, handler.ConfBinder))
	}
	root.Run("0.0.0.0:8084")
}
