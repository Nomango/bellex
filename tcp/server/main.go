// Copyright (C) 2018 Nomango - All Rights Reserved

package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"github.com/nomango/bellex/services/tcp"
	"github.com/nomango/bellex/services/tcp/types"
)

func main() {
	tcpServer, err := tcp.NewServer()
	if err != nil {
		log.Fatalln("Start server failed: ", err)
	}

	defer tcpServer.Close()

	log.Println("TCP server is running on", tcpServer.Addr())

	// start to accept connections
	for {
		conn, err := tcpServer.Accept()
		if err != nil {
			log.Println("Accept TCP connection failed,", err)
			continue
		}
		log.Println("Accept TCP connection from", conn.RemoteAddr().String())

		// handle conn in goroutine
		go tcpServer.Handle(conn, handlePacket)
	}
}

func handlePacket(packet *types.Packet, conn net.Conn) {
	// var verifyPacket types.VerifyPacket
	// if err := json.Unmarshal(packet.PacketContent, &verifyPacket); err != nil {
	// 	log.Println("Unmarshal json data failed", conn.RemoteAddr().String(), err)
	// 	return
	// }

	log.Println("Request from", conn.RemoteAddr().String())

	switch packet.Type {
	case types.PacketTypeGetTime:
		timestamp := strconv.FormatInt(time.Now().Unix(), 10) + "\n"
		conn.Write([]byte(timestamp))
		return
	case types.PacketTypeChangeMode:
		conn.Write([]byte("PacketTypeChangeMode data has received\n"))
		return
	}
}
