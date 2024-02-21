package client

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

var (
	HC *HttpClient
)

type HttpClient struct {
	Client *http.Client
}

func init() {
	caCert, err := os.ReadFile("conf/ca.crt")
	if err != nil {
		log.Fatalf("Reading CA certificate: %s", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// 读取客户端证书和私钥
	clientCert, err := tls.LoadX509KeyPair("conf/client.crt", "conf/client.key")
	if err != nil {
		log.Fatalf("Loading client key pair: %s", err)
	}

	// 创建TLS配置
	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{clientCert},
	}
	HC = &HttpClient{
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
	}
}

// func InitClient(tlsConfig *tls.Config) {
// 	HC = &HttpClient{
// 		Client: &http.Client{
// 			Transport: &http.Transport{
// 				TLSClientConfig: tlsConfig,
// 			},
// 		},
// 	}
// }

// RequestWithHead is a shortcut of func(hc *HttpClient) RequestWithHead(){}
func RequestWithHead(method, host, uri string, header http.Header, param interface{}) (interface{}, error) {
	return HC.RequestWithHead(method, host, uri, header, param)
}

func getBytes(data interface{}) (result []byte, err error) {
	if data == nil {
		return nil, nil
	}

	switch v := data.(type) {
	case string:
		result = []byte(v)
	case []byte:
		result = v
	default:
		result, err = json.Marshal(data)
	}
	return
}

func do(method, url string, header http.Header, param interface{}) (interface{}, error) {
	var err error
	requestBody, err := getBytes(param)
	if err != nil {
		logutil.Error("ready to post to [%v], data: [%+v]", url, param)
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		logutil.Error("creat request failed")
		return nil, err
	}
	req.Header = header
	resp, err := HC.Client.Do(req)
	if err != nil {
		logutil.Error("client do request failed")
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logutil.Error("read resp.Body failed, err: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		logutil.Error("unmarshal response to struct failed, err: %v", err)
	}
	return result, nil
}
func (hc *HttpClient) RequestWithHead(method, host, uri string, header http.Header, param interface{}) (interface{}, error) {
	logutil.Info("ready to post to host: %v, uri: %v", host, uri)
	url := host + uri
	if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
		url = fmt.Sprintf("https://%s", url)
	}
	return do(method, url, header, param)
}
