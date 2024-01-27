package midware

import (
	"fmt"
	errcode "github.com/OrigamiWang/msd/micro/const"
	"github.com/OrigamiWang/msd/micro/model"
	"github.com/OrigamiWang/msd/micro/model/errx"
	"github.com/OrigamiWang/msd/micro/util"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

type PostHandlerFunc func(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX)

type HandlerReqBinder func() (req interface{})

func PostHandler(handlerfunc PostHandlerFunc, binder ...HandlerReqBinder) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		funcName := runtime.FuncForPC(reflect.ValueOf(handlerfunc).Pointer()).Name()

		var resp interface{}
		var err errx.ErrX
		if len(binder) > 0 {
			req := binder[0]()
			bindErr := c.BindJSON(req)
			if bindErr != nil {
				logutil.Error("PostHandler. handler begin, Bind json failed, funcName: %v, error: %v", funcName, bindErr)
				c.JSON(http.StatusOK, &model.Response{Code: errcode.WrongArgs, Ts: fmt.Sprintf("%v", time.Now().Unix()), Msg: "Wrong argument"})
				return
			}
			logutil.Info("PostHandler. handler begin, funcName: %v, req: %s", funcName, util.ReflectToString(req))
			resp, err = handlerfunc(c, req)
		} else {
			logutil.Info("PostHandler. handler begin, funcName: %v, req: <nil>", funcName)
			resp, err = handlerfunc(c, nil)
		}
		if err == nil {
			c.Render(http.StatusOK, render.JSON{Data: &model.Response{Code: errcode.Success, Ts: fmt.Sprintf("%v", time.Now().Unix()), Msg: "Success", Data: resp}})
			logutil.Info("PostHandler. handler finish, funcName: %v, code: %v, msg: %v. elapse: %v", funcName, errcode.Success, "success", time.Since(start))
			logutil.Debug("PostHandler. handler finish, funcName: %v, resp: %+v", funcName, util.ReflectToString(resp))
		} else {
			c.JSON(http.StatusOK, &model.Response{Code: err.Code(), Ts: fmt.Sprintf("%v", time.Now().Unix()), Msg: err.Error(), Data: resp})
			logutil.Info("PostHandler. handler finish, funcName: %v, code: %v, msg: %v. elapse: %v", funcName, err.Code(), err.Error(), time.Since(start))
			logutil.Debug("PostHandler. handler finish, funcName: %v, resp: %+v", funcName, util.ReflectToString(resp))
		}
		c.Next()
	}
}
