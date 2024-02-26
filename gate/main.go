package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/OrigamiWang/msd/gate/biz"
	"github.com/OrigamiWang/msd/gate/biz/lb"
	tls2 "github.com/OrigamiWang/msd/micro/auth/tls"
	"github.com/OrigamiWang/msd/micro/framework"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/OrigamiWang/msd/register-center/model/dto"
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
	svcMap, err := biz.GetLiveSvc()
	if err != nil {
		logutil.Error("fail to get svc map, err: %v", err)
	}
	fmt.Println("svcMap: ", svcMap)
	lb := lb.NewLoadBalancer()
	for k, v := range svcMap {
		fmt.Println(k)
		instJson := v.(string)
		instConf := &dto.InstanceConf{}
		json.Unmarshal([]byte(instJson), instConf)
		fmt.Println(instConf)
		url := fmt.Sprintf("https://%v:%v", instConf.Ip, instConf.Port)
		svc_name := strings.Split(k, "_")[0]
		lb.AddService(svc_name, url)
	}
	for k, _ := range lb.GetAllBalancer() {
		prefix := fmt.Sprintf("/%s", k)
		biz.Proxy(lb, prefix, public, client)
	}

	if err := root.Run(":8848"); err != nil {
		logutil.Error("open gateway failed, err: %v", err)
	}
}
