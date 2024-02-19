package biz

import (
	"github.com/OrigamiWang/msd/manage/dal"
	"github.com/OrigamiWang/msd/manage/model/dto"
	"github.com/OrigamiWang/msd/micro/auth"
	"github.com/OrigamiWang/msd/micro/auth/crypto"
	"github.com/OrigamiWang/msd/micro/const/errcode"
	"github.com/OrigamiWang/msd/micro/model/errx"
)

func DoLogin(req *dto.LoginReq) (resp interface{}, err errx.ErrX) {
	err = nil
	name := req.Name
	pswd := req.Pswd
	encodedPswd := crypto.Md5Encode(pswd)
	userDao, e := dal.Login(name, encodedPswd)
	if e != nil {
		return nil, errx.New(errcode.Success, "name or password incorrect")
	}
	// return jwt

	jwt := auth.Authorize(userDao.ID, userDao.Name)
	resp = map[string]interface{}{
		"uid":   userDao.ID,
		"uname": userDao.Name,
		"jwt":   jwt,
	}
	return
}
