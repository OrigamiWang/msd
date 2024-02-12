package facade

import (
	"net/http"

	httpmethod "github.com/OrigamiWang/msd/micro/const/http"
	"github.com/OrigamiWang/msd/micro/framework/client"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

type IFUser interface {
	GetUserList() (interface{}, error)
}

type UserManageFacade struct {
}

func (this *UserManageFacade) GetUserList() (interface{}, error) {
	resp, err := client.RequestWithHead(httpmethod.GET, "localhost:8081", "/user", http.Header{}, nil)
	if err != nil {
		logutil.Error("request with head failed, err: %v", err)
		return nil, err
	} else {
		return resp, nil
	}
}
