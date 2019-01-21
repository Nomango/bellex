// Copyright (C) 2018 Nomango - All Rights Reserved

package server

import (
	"errors"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/nomango/bellex/services/ntp"
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

// HandlePacket handle request packets
func HandlePacket(packet *types.Packet, conn net.Conn) {

	if !Verify(packet) {
		write("Permission denied", conn)
		return
	}

	switch packet.Type {
	case types.PacketTypeRequestTime:
		if now, err := requestNTP(); err != nil {
			write(err.Error(), conn)
		} else {
			write(strconv.FormatInt(now, 10), conn)
		}
	case types.PacketTypeChangeMode:
		write("PacketTypeChangeMode data has received", conn)
	}
}

func requestNTP() (int64, error) {
	size := len(ntp.Servers)
	signals := make(chan time.Time, size)

	for _, host := range ntp.Servers {
		go func(host string) {
			if now, err := ntp.SendRequest(host); err == nil {
				signals <- now
			} else {
				signals <- time.Time{}
			}
		}(host)
	}

	var sum int64
	for range make([]int, size) {
		result := <-signals
		if result.IsZero() {
			return 0, errors.New("Request NTP failed")
		}
		sum += result.Unix()
	}
	return sum / int64(size), nil
}

func write(data string, conn net.Conn) error {
	bytes := append([]byte(data), byte(0))
	if _, err := conn.Write(bytes); err != nil {
		return err
	}
	return nil
}
