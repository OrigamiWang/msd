package tls

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

var (
	TlsClientConfig = &tls.Config{}
	TlsServerConfig = &tls.Config{}
)

func init() {
	var err error
	TlsClientConfig, err = initTlsClientConfig()
	if err != nil {
		logutil.Error("init cls client config failed, err: %v", err)
		panic(err.Error())
	}
	TlsServerConfig, err = initTlsServerConfig()
	if err != nil {
		logutil.Error("init cls server config failed, err: %v", err)
		panic(err.Error())
	}
}
func initCertAndPool() (*tls.Certificate, *x509.CertPool, error) {
	caCert, err := os.ReadFile("conf/ca.crt")
	if err != nil {
		logutil.Error("read ca.crt error: %v", err)
		return nil, nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	clientCert, err := tls.LoadX509KeyPair("conf/client.crt", "conf/client.key")
	if err != nil {
		logutil.Error("read client.crt or client.key error: %v", err)
		return nil, nil, err
	}
	return &clientCert, caCertPool, nil
}

func initTlsClientConfig() (*tls.Config, error) {
	cert, pool, err := initCertAndPool()
	if err != nil {
		logutil.Error("init certificate and pool failed err: %v", err)
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*cert},
		RootCAs:      pool,
	}
	return tlsConfig, nil
}
func initTlsServerConfig() (*tls.Config, error) {
	cert, pool, err := initCertAndPool()
	if err != nil {
		logutil.Error("init certificate and pool failed err: %v", err)
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*cert},
		RootCAs:      pool,
	}
	return tlsConfig, nil
}
