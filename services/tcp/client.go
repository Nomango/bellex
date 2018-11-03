// Copyright (C) 2018 Nomango - All Rights Reserved

package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"

	"github.com/nomango/bellex/services/tcp/types"
)

const (
	serverIP = "139.199.207.247:" + serverPort
)

func packSendData(sendBytes []byte) []byte {
	size := len(sendBytes) + 6
	result := make([]byte, size)
	result[0] = 0xFF
	result[1] = 0xFF
	result[2] = byte(uint16(len(sendBytes)) >> 8)
	result[3] = byte(uint16(len(sendBytes)) & 0xFF)

	copy(result[4:], sendBytes)

	result[size-2] = 0xFF
	result[size-1] = 0xFE
	return result
}

func getRandString() string {
	length := rand.Intn(50)
	strBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		strBytes[i] = byte(rand.Intn(26) + 97)
	}
	return string(strBytes)
}

// Client tcp client
type Client struct {
	conn *net.TCPConn
	addr *net.TCPAddr
	Stop chan struct{}
}

// NewClient returns a new tcp cliet
func NewClient() (*Client, error) {
	addr, err := net.ResolveTCPAddr("tcp", serverIP)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
		addr: addr,
		Stop: make(chan struct{}),
	}, nil
}

// Receive get server response
func (c *Client) Receive() {
	reader := bufio.NewReader(c.conn)
	for {
		response, err := reader.ReadString('\n')
		if err != nil {
			close(c.Stop)
			break
		}
		fmt.Print("Server response: ", response)
	}
}

// RequestServerTime send 'GetTime' request
func (c *Client) RequestServerTime() {

	packet := types.Packet{
		Type: types.PacketTypeGetTime,
		Data: make([]byte, 0),
	}
	packetData, err := json.Marshal(packet)
	if err != nil {
		fmt.Println("Marshal json data failed,", err)
		return
	}

	data := packSendData(packetData)
	fmt.Println("Send data:", data)

	if _, err := c.conn.Write(data); err != nil {
		fmt.Println("Send request failed,", err)
		return
	}
}
