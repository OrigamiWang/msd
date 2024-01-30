package handler

import (
	"github.com/OrigamiWang/msd/auth/cli"
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
)

// 测试cli接口
func TestHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	result, e := cli.Manage.GetUserList()
	if e != nil {
		logutil.Error("err: %v", err)
	} else {
		logutil.Info("result: %v", result)
	}
	return nil, nil
}
