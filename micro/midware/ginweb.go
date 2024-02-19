package midware

import (
	"net/http"

	"github.com/OrigamiWang/msd/micro/const/apikey"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"github.com/gin-gonic/gin"
)

func CheckGateApiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey != apikey.GATE_API_KEY {
			logutil.Warn("invalid api key")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid api key"})
			c.Abort() // prevent to call the function next
			return
		}
		c.Next()
	}
}

func AddGateApiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Add("X-API-Key", apikey.GATE_API_KEY)
		c.Next()
	}
}
