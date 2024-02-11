package handler

import (
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
)

// 测试cli接口
func TestHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	logutil.Info("call test handler...")
	return nil, nil
}
