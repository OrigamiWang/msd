package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"

	"github.com/OrigamiWang/msd/conf-center/dal"
	"github.com/OrigamiWang/msd/conf-center/handler"
	"github.com/OrigamiWang/msd/micro/confparser"
	"github.com/OrigamiWang/msd/micro/framework"
	mw "github.com/OrigamiWang/msd/micro/midware"
	"github.com/OrigamiWang/msd/micro/model/dao"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

func init() {
	// conf-center get config by conf.yml
	// other svc get conf by conf-center
	confparser.Conf = confparser.LoadConf()
	dao.InitDb()
	dal.InitConn()
}

func main() {

	caCert, err := os.ReadFile("conf/ca.crt")
	if err != nil {
		logutil.Error("load ca.crt error: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	confCenterCert, err := tls.LoadX509KeyPair("conf/conf-center.crt", "conf/conf-center.key")
	if err != nil {
		logutil.Error("load conf-center.crt error: %v", err)
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{confCenterCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
	}

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

	srv := &http.Server{
		Addr:      ":8084",
		Handler:   root,
		TLSConfig: tlsConfig,
	}
	if err := srv.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
		log.Fatalf("启动服务器失败: %v", err)
	}
}

// func main() {
// 	root := framework.NewGinWeb()
// 	r := root.Group("/")
// 	d := root.Group("/debug")
// 	{
// 		d.GET("/test", mw.PostHandler(handler.TestHandler))
// 	}
// 	{
// 		r.GET("/config/:name", handler.GetConfigHandler)
// 		r.PUT("/config/:name", mw.PostHandler(handler.UpdateConfigHandler, handler.ConfBinder))
// 	}
// 	root.Run("0.0.0.0:8084")
// }
