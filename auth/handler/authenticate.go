package handler

import (
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
)

func AuthenticateHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	logutil.Info("authenticate...")
	return nil, nil
}
