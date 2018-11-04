// Copyright (C) 2018 Nomango - All Rights Reserved

package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nomango/bellex/services/api"
	"github.com/nomango/bellex/services/tcp"
	"github.com/nomango/bellex/services/tcp/types"
)

func main() {

	go startTCPServer()

	// Windows console cannot display color correctly, so disable it
	gin.DisableConsoleColor()

	engine := gin.Default()
	api.SetupRouter(engine)

	// Listen and Server in 0.0.0.0:8080
	engine.Run(":8080")
}

func startTCPServer() {
	tcpServer, err := tcp.NewServer()
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
		go tcpServer.Handle(conn, handlePacket)
	}
}

func handlePacket(packet *types.Packet, conn net.Conn) {

	now := time.Now().Format("| 2006-01-02 15:04:05 |")
	log.Println("[Bellex]", now, "Request from", conn.RemoteAddr().String())

	switch packet.Type {
	case types.PacketTypeGetTime:
		stamp := strconv.FormatInt(time.Now().Unix(), 10) + "\n"
		conn.Write([]byte(stamp))
		return
	case types.PacketTypeChangeMode:
		conn.Write([]byte("PacketTypeChangeMode data has received\n"))
		return
	}
}
