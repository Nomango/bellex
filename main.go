// Copyright (C) 2018 Nomango - All Rights Reserved

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nomango/bellex/services/api"
	"github.com/nomango/bellex/tcp/server"
)

func main() {

	// start a tcp server
	go server.Start()

	// Windows PowerShell cannot display color correctly, so disable it
	gin.DisableConsoleColor()

	engine := gin.Default()
	api.SetupRouter(engine)

	// Listen and Server in 0.0.0.0:8080
	engine.Run(":8080")
}
