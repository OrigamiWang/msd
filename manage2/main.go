package main

import (
	"github.com/OrigamiWang/msd/manage2/cli"
	"github.com/OrigamiWang/msd/manage2/dal"
	"github.com/OrigamiWang/msd/manage2/handler"
	"github.com/OrigamiWang/msd/micro/confparser"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
	"github.com/OrigamiWang/msd/micro/model/dao"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/mitchellh/mapstructure"
)

func init() {
	resp, err := cli.Conf.GetConf("manage2")
	if err != nil {
		logutil.Error("get conf failed, err: %v", err)
	}
	m := resp.(map[string]interface{})
	logutil.Info(m)
	var conf *confparser.Config
	err = mapstructure.Decode(m, &conf)
	if err != nil {
		logutil.Error("marshal json failed, err: %v", err)
		panic("marshal json failed")
	}
	confparser.Conf = conf
	dao.InitDb()
	dal.InitConn()
}
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
	root.Run("localhost:8082")
}
