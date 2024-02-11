package handler

import (
	"github.com/OrigamiWang/msd/auth/cli"
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
)

// 测试cli接口
func UserHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	result, e := cli.Manage.GetUserList()
	if e != nil {
		logutil.Error("err: %v", err)
	} else {
		res := result.(map[string]interface{})
		// json默认解码为float64
		if res["code"].(float64) == 0 {
			logutil.Info(res["data"])
		} else {
			logutil.Warn("something wrong")
		}
	}
	return nil, nil
}
