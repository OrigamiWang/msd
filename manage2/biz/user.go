package biz

import (
	"github.com/OrigamiWang/msd/manage/dal"
	"github.com/OrigamiWang/msd/manage/model/dto"
	"github.com/OrigamiWang/msd/micro/auth/crypto"
	"github.com/OrigamiWang/msd/micro/const/errcode"
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
)

func GetUserById(uid string) (resp *dto.UserResp, err errx.ErrX) {
	userDao, e := dal.GetUserById(uid)
	if e != nil {
		logutil.Error("mysql. get user id failed, err: %v", e)
		return nil, errx.New(errcode.MysqlErr, e.Error())
	}
	resp = &dto.UserResp{
		ID:   userDao.ID,
		Name: userDao.Name,
		Age:  userDao.Age,
		Sex:  userDao.Sex,
	}
	return
}

func GetUserList() (resp *[]dto.UserResp, err errx.ErrX) {
	userDaoArr, e := dal.GetAllUser()
	if e != nil {
		logutil.Error("mysql. get all user failed, err: %v", e)
		return nil, errx.New(errcode.MysqlErr, e.Error())
	}
	userRespArr := []dto.UserResp{}
	for _, userDao := range *userDaoArr {
		userRespArr = append(userRespArr, dto.UserResp{
			ID:   userDao.ID,
			Name: userDao.Name,
			Age:  userDao.Age,
			Sex:  userDao.Sex,
		})
	}
	resp = &userRespArr
	return
}

func UpdateUser(uid string, userReq *dto.UserReq) (resp *dto.UserResp, err errx.ErrX) {
	userReq.Pswd = crypto.Md5Encode(userReq.Pswd)
	userDao, e := dal.UpdateUser(uid, userReq)
	if e != nil {
		logutil.Error("mysql. update user failed, err: %v", e)
		return nil, errx.New(errcode.MysqlErr, e.Error())
	}
	resp = &dto.UserResp{
		ID:   userDao.ID,
		Name: userDao.Name,
		Age:  userDao.Age,
		Sex:  userDao.Sex,
	}
	return
}

func AddUser(userReq *dto.UserReq) (resp *dto.UserResp, err errx.ErrX) {
	if userReq.Age <= 0 || userReq.Sex == "" {
		return nil, errx.New(errcode.WrongArgs, "wrong args")
	}
	user, e := dal.GetUserByName(userReq.Name)
	if user != nil && e == nil {
		// name can not be repeated
		return nil, errx.New(errcode.WrongArgs, "the name can not be repeated")
	}
	userReq.Pswd = crypto.Md5Encode(userReq.Pswd)
	userDao, e := dal.AddUser(userReq)
	if e != nil {
		logutil.Error("mysql. add user failed, err: %v", e)
		return nil, errx.New(errcode.MysqlErr, e.Error())
	}
	resp = &dto.UserResp{
		ID:   userDao.ID,
		Name: userDao.Name,
		Age:  userDao.Age,
		Sex:  userDao.Sex,
	}
	return
}
