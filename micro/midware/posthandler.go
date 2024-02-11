package midware

import (
	"fmt"
	"github.com/OrigamiWang/msd/micro/auth"
	"github.com/OrigamiWang/msd/micro/const/errcode"
	"github.com/OrigamiWang/msd/micro/model"
	"github.com/OrigamiWang/msd/micro/model/errx"
	"github.com/OrigamiWang/msd/micro/util"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
	"reflect"
	"runtime"
	"strings"
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

func PostHandlerWithJwt(handlerfunc PostHandlerFunc, binder ...HandlerReqBinder) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		funcName := runtime.FuncForPC(reflect.ValueOf(handlerfunc).Pointer()).Name()

		// check the jwt token
		// Authorization: 'Bearer JwtToken'
		authorization := c.GetHeader("Authorization")
		m, e := checkJwt(authorization)
		if e != nil {
			logutil.Error("check jwt failed, err: %v", e)
			c.JSON(http.StatusOK, &model.Response{Code: errcode.WrongJwt, Ts: fmt.Sprintf("%v", time.Now().Unix()), Msg: "Wrong jwt"})
			return
		}
		uid := m["uid"].(int)
		uname := m["uname"].(string)
		logutil.Info("check jwt success, uid: %v, uname: %v", uid, uname)

		var resp interface{}
		var err errx.ErrX
		if len(binder) > 0 {

			// get request body
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

func checkJwt(authorization string) (map[string]interface{}, error) {
	if authorization == "" {
		logutil.Error("the authorization is nil")
		return nil, fmt.Errorf("the authorization is nil")
	}
	arr := strings.Split(authorization, " ")
	// invalid
	if arr == nil || len(arr) != 2 || arr[0] != "Bearer" {
		logutil.Error("invalid request")
		return nil, fmt.Errorf("invalid request")
	}
	jwtToken := arr[1]
	uid, uname, err := auth.Authenticate(jwtToken)
	if err != nil {
		logutil.Error("cli. jwt token authenticate failed, err: %v", err)
		return nil, err
	}
	m := make(map[string]interface{})
	m["uid"] = uid
	m["uname"] = uname
	return m, nil
}
