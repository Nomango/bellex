// Copyright (C) 2018 Nomango - All Rights Reserved

package tcp

import (
	"bufio"
	"errors"
	"log"
	"net"
)

// Server tcp server
type Server struct {
	listener *net.TCPListener
	addr     *net.TCPAddr
}

// getLocalIP get local ip address
func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("Local IP not found")
}

// NewServer returns a new tcp server
func NewServer(port string) (*Server, error) {
	localIP, err := getLocalIP()
	if err != nil {
		return nil, err
	}
	addr, err := net.ResolveTCPAddr("tcp", localIP+":"+port)
	if err != nil {
		return nil, err
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener: listener,
		addr:     addr,
	}, nil
}

// Accept waits for the next call and returns a generic Conn
func (ts *Server) Accept() (net.Conn, error) {
	return ts.listener.Accept()
}

// Close stop listening on the TCP address
func (ts *Server) Close() {
	ts.listener.Close()
}

// Addr returns the TCP address which is listening
func (ts *Server) Addr() string {
	return ts.addr.String()
}

// Handle recives data from client and send response
// Data format: 0xFF|0xFF|len(high)|len(low)|Data|0xFF|0xFE. '0xFF' is preamble code
func (ts *Server) Handle(conn net.Conn, handler func([]byte, net.Conn)) {
	// close connection before exit
	defer conn.Close()

	var (
		dataSize     uint8
		dataCursor   uint8
		recvBuffer   []byte
		bufferReader = bufio.NewReader(conn)
	)

	// state machine
	state := 0
	for {
		recvByte, err := bufferReader.ReadByte()
		if err != nil {
			log.Println("Connection " + conn.RemoteAddr().String() + " is closed")
			return
		}

		switch state {
		case 0:
			if recvByte == 0xFF {
				state = 1
				recvBuffer = nil
				dataSize = 0
			} else {
				state = 0
			}
		case 1:
			if recvByte == 0xFF {
				state = 2
			} else {
				state = 0
			}
		case 2:
			dataSize = uint8(recvByte)
			recvBuffer = make([]byte, dataSize)
			dataCursor = 0
			state = 3
		case 3:
			recvBuffer[dataCursor] = recvByte
			dataCursor++
			if dataCursor == dataSize {
				state = 4
			}
		case 4:
			if recvByte == 0xFF {
				state = 5
			} else {
				state = 0
			}
		case 5:
			if recvByte == 0xFE {
				handler(recvBuffer, conn)
			}
			// state machine is ready. read next packet.
			state = 0
		}
	}
}