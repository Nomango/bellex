package main

import (
	"log"
	"net/http"

	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/settings"
	"github.com/nomango/bellex/services/tcp"
	"github.com/nomango/bellex/services/tcp/tcprpc"
)

func main() {
	log.SetPrefix("[bellex] ")

	settings.Setup()
	models.Setup()

	listenRPC()
	listenTCP()
}

func listenRPC() {
	rpcServer, err := tcprpc.NewServer()
	if err != nil {
		log.Fatalln("Start RPC server failed:", err)
	}
	log.Println("RPC server is running on", rpcServer.Addr())

	go http.Serve(rpcServer, nil)
}

func listenTCP() {
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
