package handler

import (
	"net/http"

	"github.com/OrigamiWang/msd/micro/confparser"
	httpmethod "github.com/OrigamiWang/msd/micro/const/http"
	"github.com/OrigamiWang/msd/micro/framework/client"
	"github.com/OrigamiWang/msd/micro/model/errx"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	return testPostWithHead(), nil
}

func GetConfExtHandler(c *gin.Context, req interface{}) (resp interface{}, err errx.ErrX) {
	key := c.Param("key")
	logutil.Info("key: %s", key)
	ext := confparser.ExtString(key)
	c.JSON(http.StatusOK, gin.H{"ext": ext})
	return nil, nil
}

func testPostWithHead() interface{} {
	header := http.Header{}
	header.Set("Content-Type", "application/json")
	result, err := client.RequestWithHead(httpmethod.GET, "localhost:8081", "/user/1", header, nil)
	if err != nil {
		logutil.Error("request with head failed, err: %v", err)
		return nil
	}
	return result
}
