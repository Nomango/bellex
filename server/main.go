package main

import (
	"log"

	_ "github.com/nomango/bellex/server/routers"
	"github.com/nomango/bellex/server/services/tcp"

	"github.com/astaxie/beego"
)

func main() {
	// start a tcp server
	go startTCPServer()

	beego.SetViewsPath("views")
	beego.SetStaticPath("/static", "static")

	beego.Run()
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
