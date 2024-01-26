package midware

import (
	"github.com/OrigamiWang/msd/micro/model/errx"
	"github.com/gin-gonic/gin"
)

type PostHandlerFunc func(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX)

type HandlerReqBinder func() (req interface{})

func PostHandler(handlerfunc PostHandlerFunc, binder ...HandlerReqBinder) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
