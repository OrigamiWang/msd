package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/OrigamiWang/msd/conf-center/handler"
	"github.com/OrigamiWang/msd/micro/auth/tls"
	"github.com/OrigamiWang/msd/micro/const/svc"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
	"github.com/OrigamiWang/msd/register-center/biz"
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

	// set config
	var ip string
	var port int
	var instanceId int
	flag.StringVar(&ip, "ip", "127.0.0.1", "ip")
	flag.IntVar(&port, "port", 8849, "port")
	flag.IntVar(&instanceId, "instance_id", 0, "instance id")
	flag.Parse()
	hostMap := make(map[string]interface{})
	hostMap["ip"] = ip
	hostMap["port"] = port
	hostMap["instance_id"] = instanceId
	jsonBytes, _ := json.Marshal(&hostMap)
	biz.BeatHeartBeat(svc.CONF_CENTER, string(jsonBytes))
	addr := fmt.Sprintf(":%v", port)

	tlsConfig := tls.TlsServerConfig
	srv := &http.Server{
		Addr:      addr,
		Handler:   root,
		TLSConfig: tlsConfig,
	}
	if err := srv.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
