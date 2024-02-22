package facade

import (
	"fmt"
	"net/http"

	httpmethod "github.com/OrigamiWang/msd/micro/const/http"
	"github.com/OrigamiWang/msd/micro/framework/client"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

type IFRegisterCenter interface {
}

type RegisterCenterFacade struct {
}

func (facade *RegisterCenterFacade) GetAllSvc() (interface{}, error) {
	resp, err := client.RequestWithHead(httpmethod.GET, "localhost:8000", "/register", http.Header{}, nil)
	if err != nil {
		logutil.Error("request with head failed, err: %v", err)
		return nil, err
	} else {
		return resp, nil
	}
}
func (facade *RegisterCenterFacade) GetSvcByName(name string) (interface{}, error) {
	uri := fmt.Sprintf("/register/%s", name)
	resp, err := client.RequestWithHead(httpmethod.GET, "localhost:8000", uri, http.Header{}, nil)
	if err != nil {
		logutil.Error("request with head failed, err: %v", err)
		return nil, err
	} else {
		return resp, nil
	}
}
