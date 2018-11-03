package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter(engine *gin.Engine) {
	// Ping test
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
