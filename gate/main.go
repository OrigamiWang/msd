package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OrigamiWang/msd/gate/biz"
	"github.com/OrigamiWang/msd/gate/cli"
	tls2 "github.com/OrigamiWang/msd/micro/auth/tls"
	"github.com/OrigamiWang/msd/micro/framework"
	"github.com/OrigamiWang/msd/micro/model"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/mitchellh/mapstructure"
)

func main() {
	root := framework.NewGinWeb()
	r := root.Group("/api")

	public := r.Group("/public")
	// auth := r.Group("/auth")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tls2.TlsServerConfig,
		},
	}

	// get service conf
	resp, err := cli.Conf.GetAllSvc()
	if err != nil {
		logutil.Error("get all register config failed, err: %v", err)
	}
	m := resp.(map[string]interface{})
	var data model.Response
	mapstructure.Decode(m, &data)
	arr := data.Data.([]interface{})
	for _, raw := range arr {
		item, ok := raw.(map[string]interface{})
		if !ok {
			logutil.Warn("Type assertion failed")
			break
		}
		if item["name"] == "gate" {
			continue
		}
		var registerConfig map[string]interface{}
		err := json.Unmarshal([]byte(item["config"].(string)), &registerConfig)
		if err != nil {
			logutil.Error("decode config failed, err: %v", err)
			break
		}
		logutil.Info("name: %v, ip: %v, port: %v", item["name"], registerConfig["ip"], registerConfig["port"])

		url := fmt.Sprintf("https://%v:%v", registerConfig["ip"], registerConfig["port"])
		prefix := fmt.Sprintf("/%s", item["name"])

		// proxy
		biz.Proxy(url, prefix, public, client)
	}

	if err := root.Run(":8848"); err != nil {
		logutil.Error("open gateway failed, err: %v", err)
	}
}
