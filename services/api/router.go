// Copyright (C) 2018 Nomango - All Rights Reserved

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		// Ping test
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}
}
