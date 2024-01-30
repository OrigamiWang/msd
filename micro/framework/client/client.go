package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
	HC = &HttpClient{
		Client: &http.Client{},
	}
}

// PostWithHead is a shortcut of func(hc *HttpClient) PostWithHead(){}
func RequestWithHead(method, host, uri string, header http.Header, param interface{}, resp *http.Response) error {
	return HC.RequestWithHead(method, host, uri, header, param, resp)
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

func do(method, url string, header http.Header, param interface{}, resp *http.Response) error {
	var err error
	requestBody, err := getBytes(param)
	if err != nil {
		logutil.Error("ready to post to [%v], data: [%+v]", url, param)
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		logutil.Error("creat request failed")
		return err
	}
	req.Header = header
	resp, err = HC.Client.Do(req)
	if err != nil {
		logutil.Error("client do request failed")
		return err
	}
	return nil
}
func (hc *HttpClient) RequestWithHead(method, host, uri string, header http.Header, param interface{}, resp *http.Response) error {
	logutil.Info("ready to post to host: %v, uri: %v", host, uri)
	url := host + uri
	if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
		url = fmt.Sprintf("http://%s", url)
	}
	return do(method, url, header, param, resp)
}
