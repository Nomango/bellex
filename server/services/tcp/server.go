// Copyright (C) 2018 Nomango - All Rights Reserved

package tcp

import (
	"bufio"
	"errors"
	"log"
	"net"
	"time"

	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/settings"
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
	var (
		addr *net.TCPAddr
		err  error
	)

	if settings.Mode == settings.ModeDevelope {
		addr, err = net.ResolveTCPAddr("tcp", "127.0.0.1:"+port)
	} else {
		var localIP string
		localIP, err = getLocalIP()
		if err != nil {
			return nil, err
		}
		addr, err = net.ResolveTCPAddr("tcp", localIP+":"+port)
	}

	if err != nil {
		return nil, err
	}

	var listener *net.TCPListener
	if listener, err = net.ListenTCP("tcp", addr); err != nil {
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
func (ts *Server) Handle(conn net.Conn, handler func([]byte, net.Conn, chan<- []byte)) {
	// close connection before exit
	defer conn.Close()

	// Wait for response
	outputCh := make(chan []byte)
	endCh := make(chan struct{})
	go syncWriter(conn, outputCh, endCh)

	// state machine
	var (
		state        int
		dataSize     uint8
		dataCursor   uint8
		recvByte     byte
		recvBuffer   []byte
		bufferReader = bufio.NewReader(conn)

		inputCh = make(chan byte)
		errCh   = make(chan error)
	)

tcpLoop:
	for {
		go syncReadByte(bufferReader, inputCh, errCh)

		select {
		case recvByte = <-inputCh:
			break
		case <-errCh:
			// connection closed
			log.Println("[Bellex] Connection " + conn.RemoteAddr().String() + " is closed")
			break tcpLoop
		case <-time.After(45 * time.Second):
			// connection timeout
			log.Println("[Bellex] Connection " + conn.RemoteAddr().String() + " timeout")
			break tcpLoop
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
				handler(recvBuffer, conn, outputCh)
			}
			// state machine is ready. read next packet.
			state = 0
		}
	}

	// send end message
	endCh <- struct{}{}

	// remove connection
	if err := models.DeleteConnectionWithConn(conn); err != nil {
		log.Println("[Bellex] Remove connection failed", conn.RemoteAddr())
	}
}

func syncWriter(conn net.Conn, outputCh <-chan []byte, endCh <-chan struct{}) {
	for {
		select {
		case data := <-outputCh:
			if len(data) != 0 {
				if _, err := conn.Write(append(data, byte(0))); err != nil {
					log.Println("[Bellex] Bad response", conn.RemoteAddr(), err.Error())
				}
			}
		case <-endCh:
			// connection closed
			return
		}
	}
}

func syncReadByte(reader *bufio.Reader, recvCh chan<- byte, errCh chan<- error) {
	recvByte, err := reader.ReadByte()
	if err != nil {
		errCh <- err
		return
	}
	recvCh <- recvByte
}
