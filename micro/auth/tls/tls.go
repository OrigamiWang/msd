package tls

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

var (
	TlsConfig = &tls.Config{}
)

func init() {
	var err error
	TlsConfig, err = initTlsConfig()
	if err != nil {
		logutil.Error("init cls config failed, err: %v", err)
		panic(err.Error())
	}
}
func initTlsConfig() (*tls.Config, error) {
	caCert, err := os.ReadFile("conf/ca.crt")
	if err != nil {
		logutil.Error("read ca.crt error: %v", err)
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	clientCert, err := tls.LoadX509KeyPair("conf/client.crt", "conf/client.key")
	if err != nil {
		logutil.Error("read client.crt or client.key error: %v", err)
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
	}
	return tlsConfig, nil
}
