package facade

import (
	"fmt"
	httpmethod "github.com/OrigamiWang/msd/micro/const/http"
	"github.com/OrigamiWang/msd/micro/framework/client"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"net/http"
)

type IFConf interface {
	GetConf(svcName string) (interface{}, error)
}

type ConfCenterFacade struct {
}

func (this *ConfCenterFacade) GetConf(svcName string) (interface{}, error) {
	uri := fmt.Sprintf("/config/%s", svcName)
	resp, err := client.RequestWithHead(httpmethod.GET, "localhost:8084", uri, http.Header{}, nil)
	if err != nil {
		logutil.Error("request with head failed, err: %v", err)
		return nil, err
	} else {
		return resp, nil
	}
}
