package framework

import (
	"github.com/OrigamiWang/msd/micro/model"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	// FIXME: add customized middleware
	// g.Use(TraceExtractor())
	// g.Use(Trace())
	// g.Use(Recovery())
	// g.Use(AccessLog())
	return g
}

func NewGinWeb() *model.Ginweb {
	return &model.Ginweb{New()}
}
