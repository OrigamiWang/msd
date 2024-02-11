package handler

import (
	"github.com/OrigamiWang/msd/manage/dal"
	"github.com/OrigamiWang/msd/manage/model/dto"
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
)

func UserBinder() interface{} {
	return &dto.UserReq{}
}

func GetUserByIdHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	userid := c.Param("id")
	resp, e := dal.GetUserById(userid)
	if e != nil {
		logutil.Error("mysql. get user id failed, err: %v", e)
		return nil, nil
	}
	return resp, nil
}

func GetUserListHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	resp, e := dal.GetAllUser()
	if e != nil {
		logutil.Error("mysql. get all user failed, err: %v", e)
		return nil, nil
	}
	return resp, nil
}

func UpdateUserHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	userid := c.Param("id")
	userReq := req.(*dto.UserReq)
	resp, e := dal.UpdateUser(userid, userReq)
	if e != nil {
		logutil.Error("mysql. update user failed, err: %v", e)
		return nil, nil
	}
	return resp, nil
}

func AddUserHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	userReq := req.(*dto.UserReq)
	resp, e := dal.AddUser(userReq)
	if e != nil {
		logutil.Error("mysql. add user failed, err: %v", e)
		return nil, nil
	}
	return resp, nil
}
