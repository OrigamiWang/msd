package main

import (
	"log"
	"net/http"

	"github.com/OrigamiWang/msd/conf-center/handler"
	"github.com/OrigamiWang/msd/micro/auth/tls"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
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
	}
	tlsConfig := tls.TlsServerConfig
	srv := &http.Server{
		Addr:      ":8084",
		Handler:   root,
		TLSConfig: tlsConfig,
	}
	if err := srv.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
