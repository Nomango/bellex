// Copyright (C) 2018 Nomango - All Rights Reserved

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nomango/bellex/services/api"
	"github.com/nomango/bellex/services/tcp"
)

func main() {

	// start a tcp server
	go startTCPServer()

	// Windows PowerShell cannot display color correctly, so disable it
	gin.DisableConsoleColor()

	engine := gin.Default()
	api.SetupRouter(engine)

	// Listen and Server in 0.0.0.0:8080
	engine.Run(":8080")
}

func startTCPServer() {
	tcpServer, err := tcp.NewServer("7777")
	if err != nil {
		log.Fatalln("[Bellex] Start TCP server failed: ", err)
	}

	defer tcpServer.Close()

	log.Println("[Bellex] TCP server is running on", tcpServer.Addr())

	// start to accept connections
	for {
		conn, err := tcpServer.Accept()
		if err != nil {
			log.Println("[Bellex] Accept TCP connection failed:", err)
			continue
		}
		log.Println("[Bellex] Accept TCP connection from", conn.RemoteAddr().String())

		// handle conn in goroutine
		go tcpServer.Handle(conn, tcp.PacketHandler)
	}
}
