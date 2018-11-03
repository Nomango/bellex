package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nomango/bellex/services/api/v1/ping"
)

// SetupRouter ...
func SetupRouter(engine *gin.Engine) {
	ping.SetupRouter(engine)
}
