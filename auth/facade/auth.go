package facade

import (
	"github.com/OrigamiWang/msd/auth/model/dto"
	httpmethod "github.com/OrigamiWang/msd/micro/const/http"
	"github.com/OrigamiWang/msd/micro/framework/client"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"net/http"
)

type IFAuth interface {
	Authorize(uid int, uname string) (interface{}, error)
	Authenticate(jwtToken string) (interface{}, error)
}

type AuthFacade struct {
}

func (this *AuthFacade) Authorize(uid int, uname string) (interface{}, error) {
	param := &dto.AuthorizeReq{
		Uid:   uid,
		Uname: uname,
	}
	resp, err := client.RequestWithHead(httpmethod.POST, "localhost:8082", "/authorize", http.Header{}, param)
	if err != nil {
		logutil.Error("request with head failed, err: %v", err)
		return nil, err
	} else {
		return resp, nil
	}
}

func (this *AuthFacade) Authenticate(jwtToken string) (interface{}, error) {
	param := &dto.AuthenticateReq{
		JwtToken: jwtToken,
	}
	resp, err := client.RequestWithHead(httpmethod.POST, "localhost:8082", "/authenticate", http.Header{}, param)
	if err != nil {
		logutil.Error("request with head failed, err: %v", err)
		return nil, err
	} else {
		return resp, nil
	}
}
