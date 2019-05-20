package main

import (
	"log"

	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/settings"
	"github.com/nomango/bellex/server/services/tcp"
)

func main() {
	settings.Setup()
	models.Setup()

	tcpServer, err := tcp.NewServer(settings.TcpPort)
	if err != nil {
		log.Fatalln("[Bellex] Start TCP server failed:", err)
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
