package facade

import (
	"fmt"
	"net/http"

	httpmethod "github.com/OrigamiWang/msd/micro/const/http"
	"github.com/OrigamiWang/msd/micro/framework/client"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

type IFConf interface {
	GetConf(svcName string) (interface{}, error)
	UpdateConf(svcName string, conf interface{}) (interface{}, error)
	GetAllSvc() (interface{}, error)
	GetSvcByName(name string) (interface{}, error)
}

type ConfCenterFacade struct {
}

func (facade *ConfCenterFacade) GetConf(svcName string) (interface{}, error) {
	uri := fmt.Sprintf("/config/%s", svcName)
	resp, err := client.RequestWithHead(httpmethod.GET, "localhost:8849", uri, http.Header{}, nil)
	if err != nil {
		logutil.Error("request with head failed, err: %v", err)
		return nil, err
	} else {
		return resp, nil
	}
}

func (facade *ConfCenterFacade) UpdateConf(svcName string, conf interface{}) (interface{}, error) {
	uri := fmt.Sprintf("/config/%s", svcName)
	resp, err := client.RequestWithHead(httpmethod.PUT, "localhost:8849", uri, http.Header{}, conf)
	if err != nil {
		logutil.Error("request with head failed, err: %v", err)
		return nil, err
	} else {
		return resp, nil
	}
}

func (facade *ConfCenterFacade) GetAllSvc() (interface{}, error) {
	resp, err := client.RequestWithHead(httpmethod.GET, "localhost:8849", "/regconf", http.Header{}, nil)
	if err != nil {
		logutil.Error("request with head failed, err: %v", err)
		return nil, err
	} else {
		return resp, nil
	}
}

func (facade *ConfCenterFacade) GetSvcByName(name string) (interface{}, error) {
	uri := fmt.Sprintf("/regconf/%s", name)
	resp, err := client.RequestWithHead(httpmethod.GET, "localhost:8849", uri, http.Header{}, nil)
	if err != nil {
		logutil.Error("request with head failed, err: %v", err)
		return nil, err
	} else {
		return resp, nil
	}
}
