package main

import (
	"log"

	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/settings"
	_ "github.com/nomango/bellex/server/routers"
	"github.com/nomango/bellex/server/services/tcp"

	"github.com/astaxie/beego"
)

func main() {

	settings.Setup()
	models.Setup()

	if settings.Mode == settings.ModeProduct {
		// start a tcp server
		go startTCPServer()
	}

	beego.Run()
}

func startTCPServer() {
	tcpServer, err := tcp.NewServer("7777")
	if err != nil {
		log.Fatalln("[Mechineex] Start TCP server failed: ", err)
	}

	defer tcpServer.Close()

	log.Println("[Mechineex] TCP server is running on", tcpServer.Addr())

	// start to accept connections
	for {
		conn, err := tcpServer.Accept()
		if err != nil {
			log.Println("[Mechineex] Accept TCP connection failed:", err)
			continue
		}
		log.Println("[Mechineex] Accept TCP connection from", conn.RemoteAddr().String())

		// handle conn in goroutine
		go tcpServer.Handle(conn, tcp.PacketHandler)
	}
}
