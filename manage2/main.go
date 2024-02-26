package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OrigamiWang/msd/manage/biz/kafka"
	"github.com/OrigamiWang/msd/manage/handler"
	"github.com/OrigamiWang/msd/micro/auth/tls"
	"github.com/OrigamiWang/msd/micro/const/svc"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/OrigamiWang/msd/register-center/biz"
)

func init() {
}

func main() {
	kafka.InitConfCenterConsumer()
	root := framework.NewGinWeb()
	r := root.Group("/")
	// pprof 性能监控
	// pprof.Register(root.Engine)

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

	// get resiger config, including port
	// resp, err := cli.Conf.GetSvcByName(svc.MANAGE)
	// if err != nil {
	// 	logutil.Error("get register config failed, err: %v", err)
	// }
	// m := resp.(map[string]interface{})
	// var registerConfResp *model.Response
	// ma, err := transfer.FacadeRespToMap(m, &registerConfResp)
	// if err != nil {
	// 	logutil.Error("transfer resp failed, err: %v", err)
	// }
	// var registerConfigMap map[string]interface{}
	// json.Unmarshal([]byte(ma["config"].(string)), &registerConfigMap)

	// // start heart beat go rountine, produce msg to topic: HEART_BEAT_TOPIC per minutes, with value registerConfigMap
	// jsonBytes, err := json.Marshal(&registerConfigMap)
	// if err != nil {
	// 	logutil.Error("marshal registerConfigMap failed, err: %v", err)
	// 	return
	// }
	// biz.BeatHeartBeat(svc.MANAGE, string(jsonBytes))
	// addr := fmt.Sprintf(":%v", registerConfigMap["port"].(float64))

	// test mutiple instance
	ip := "127.0.0.1"
	port := "8082"
	hostMap := make(map[string]interface{})
	hostMap["ip"] = ip
	hostMap["port"] = port
	jsonBytes, _ := json.Marshal(&hostMap)
	biz.BeatHeartBeat(svc.MANAGE, string(jsonBytes))
	addr := fmt.Sprintf(":%v", port)

	// get tls config and run server
	tlsConfig := tls.TlsServerConfig
	srv := &http.Server{
		Addr:      addr,
		Handler:   root,
		TLSConfig: tlsConfig,
	}
	if err := srv.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
		logutil.Error("start server failed, err: %v", err)
	}

}
