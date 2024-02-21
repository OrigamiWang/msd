package main

import (
	"net/http"

	"github.com/OrigamiWang/msd/gate/biz"
	tls2 "github.com/OrigamiWang/msd/micro/auth/tls"
	"github.com/OrigamiWang/msd/micro/framework"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
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
	biz.Proxy("https://localhost:8081", "/manage", public, client)
	biz.Proxy("https://localhost:8082", "/manage2", public, client)
	biz.Proxy("https://localhost:8084", "/confcenter", public, client)
	if err := root.Run(":8848"); err != nil {
		logutil.Error("open gateway failed, err: %v", err)
	}
}
