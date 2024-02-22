package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OrigamiWang/msd/conf-center/biz"
	"github.com/OrigamiWang/msd/conf-center/handler"
	"github.com/OrigamiWang/msd/conf-center/model/dto"
	"github.com/OrigamiWang/msd/micro/auth/tls"
	"github.com/OrigamiWang/msd/micro/const/svc"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/mitchellh/mapstructure"
)

func init() {
}

func main() {

	root := framework.NewGinWeb()
	r := root.Group("/")
	d := root.Group("/debug")
	{
		d.GET("/test", mw.PostHandler(handler.TestHandler))
	}
	{
		r.GET("/config/:name", handler.GetConfigHandler)
		r.PUT("/config/:name", mw.PostHandler(handler.UpdateConfigHandler, handler.ConfBinder))
		r.GET("/regconf", mw.PostHandler(handler.GetSvcRegisterListHandler))
		r.GET("/regconf/:name", mw.PostHandler(handler.GetSvcRegisterHandler))
	}

	// get port
	resp, err := biz.GetSvcByName(svc.CONF_CENTER)
	if err != nil {
		logutil.Error("get conf center svc failed: %v", err)
	}
	var res dto.SvcRegisterResp
	mapstructure.Decode(resp, &res)
	var registerConfig map[string]interface{}
	json.Unmarshal([]byte(res.Config), &registerConfig)

	// set config
	tlsConfig := tls.TlsServerConfig
	addr := fmt.Sprintf(":%v", registerConfig["port"].(float64))
	srv := &http.Server{
		Addr:      addr,
		Handler:   root,
		TLSConfig: tlsConfig,
	}
	if err := srv.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
