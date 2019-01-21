// Copyright (C) 2018 Nomango - All Rights Reserved

package server

import (
	"log"
	"net"
	"strconv"
	"time"

	"github.com/nomango/bellex/services/tcp"
	"github.com/nomango/bellex/services/tcp/types"
)

// Start starts to listen tcp connection
func Start() {
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
		go tcpServer.Handle(conn, HandlePacket)
	}
}

// Write send string to the client
func Write(data string, conn net.Conn) error {
	bytes := append([]byte(data), byte(0))
	if _, err := conn.Write(bytes); err != nil {
		return err
	}
	return nil
}

// HandlePacket handle request packets
func HandlePacket(packet *types.Packet, conn net.Conn) {

	if !Verify(packet) {
		Write("Permission denied", conn)
		return
	}

	switch packet.Type {
	case types.PacketTypeRequestTime:
		stamp := strconv.FormatInt(time.Now().Unix(), 10)
		Write(stamp, conn)
		return
	case types.PacketTypeChangeMode:
		Write("PacketTypeChangeMode data has received", conn)
		return
	}
}
