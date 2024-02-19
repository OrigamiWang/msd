package handler

import (
	"github.com/OrigamiWang/msd/manage/biz"
	"github.com/OrigamiWang/msd/manage/model/dto"
	"github.com/OrigamiWang/msd/micro/model/errx"
	"github.com/gin-gonic/gin"
)

func UserBinder() interface{} {
	return &dto.UserReq{}
}

func GetUserByIdHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	uid := c.Param("id")
	return biz.GetUserById(uid)
}

func GetUserListHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	return biz.GetUserList()
}

func UpdateUserHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	uid := c.Param("id")
	userReq := req.(*dto.UserReq)
	return biz.UpdateUser(uid, userReq)
}

// register
func AddUserHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	userReq := req.(*dto.UserReq)
	return biz.AddUser(userReq)
}
