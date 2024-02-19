package framework

import (
	"github.com/OrigamiWang/msd/micro/midware"
	"github.com/OrigamiWang/msd/micro/model"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(midware.GateApiKeyMiddleware())
	// FIXME: add customized middleware
	// g.Use(TraceExtractor())
	// g.Use(Trace())
	// g.Use(AccessLog())
	return g
}

func NewGinWeb() *model.Ginweb {
	return &model.Ginweb{New()}
}
func NewGate() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	return g
}

// NewGateGinWeb use for gate
// this not use GateApiKeyMiddleware to check the api key
func NewGateGinWeb() *model.Ginweb {
	return &model.Ginweb{NewGate()}
}
