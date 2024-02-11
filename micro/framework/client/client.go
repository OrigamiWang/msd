package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/OrigamiWang/msd/micro/auth"
	"io"
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

// RequestWithHead is a shortcut of func(hc *HttpClient) RequestWithHead(){}
func RequestWithHead(method, host, uri string, header http.Header, param interface{}) (interface{}, error) {
	return HC.RequestWithHead(method, host, uri, header, param)
}

// RequestWithJwtTokenmethod is a shortcut of func(hc *HttpClient) RequestWithJwtTokenmethod(){}
func RequestWithJwtTokenmethod(method, host, uri string, header http.Header, param interface{}) (interface{}, error) {
	return HC.RequestWithJwtToken(method, host, uri, header, param)
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
		url = fmt.Sprintf("http://%s", url)
	}
	return do(method, url, header, param)
}

// requestHeader:
//
//	{
//		  Authorization: 'Bearer JwtToken'
//	}
func (hc *HttpClient) RequestWithJwtToken(method, host, uri string, header http.Header, param interface{}) (interface{}, error) {
	logutil.Info("ready to post to host: %v, uri: %v", host, uri)
	authorization := header.Get("Authorization")
	if authorization == "" {
		logutil.Error("the authorization is nil")
		return nil, fmt.Errorf("the authorization is nil")
	}
	arr := strings.Split(authorization, " ")
	// invalid
	if arr == nil || len(arr) != 2 || arr[0] != "Bearer" {
		logutil.Error("invalid request")
		return nil, fmt.Errorf("invalid request")
	}
	jwtToken := arr[1]
	uid, uname, err := auth.Authenticate(jwtToken)
	if err != nil {
		logutil.Error("cli. jwt token authenticate failed, err: %v", err)
		return nil, err
	}
	logutil.Info("user authenticate success, uid: %v, uname: %v", uid, uname)
	url := host + uri
	if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
		url = fmt.Sprintf("http://%s", url)
	}
	return do(method, url, header, param)
}
