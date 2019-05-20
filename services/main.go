package main

import (
	"log"

	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/settings"
	"github.com/nomango/bellex/services/tcp"
)

func main() {
	log.SetPrefix("[bellex] ")

	settings.Setup()
	models.Setup()

	tcpServer, err := tcp.NewServer(settings.TcpPort)
	if err != nil {
		log.Fatalln("Start TCP server failed:", err)
	}

	defer tcpServer.Close()

	log.Println("TCP server is running on", tcpServer.Addr())

	// start to accept connections
	for {
		conn, err := tcpServer.Accept()
		if err != nil {
			log.Println("Accept TCP connection failed:", err)
			continue
		}
		log.Println("Accept TCP connection from", conn.RemoteAddr().String())

		// handle conn in goroutine
		go tcpServer.Handle(conn, tcp.PacketHandler)
	}
}
