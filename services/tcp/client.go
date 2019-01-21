// Copyright (C) 2018 Nomango - All Rights Reserved

package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/nomango/bellex/services/tcp/types"
)

const (
	serverIP = "139.199.207.247:" + serverPort
)

// Client tcp client
type Client struct {
	conn *net.TCPConn
	addr *net.TCPAddr
	stop chan struct{}
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
	}, nil
}

// Close stop tcp connection
func (c *Client) Close() {
	c.conn.Close()
}

// Receive get server response
func (c *Client) Receive() string {
	reader := bufio.NewReader(c.conn)
	response, _ := reader.ReadString(byte(0))
	return response
}

// RequestTime send 'GetTime' request
func (c *Client) RequestTime() {

	packet := types.Packet{
		Auth: types.AuthPacket{
			ID:   "BW123",
			Code: "Xwa8pj7",
		},
		Type: types.PacketTypeRequestTime,
		Data: make([]byte, 0),
	}
	packetData, err := json.Marshal(packet)
	if err != nil {
		log.Fatalln("Marshal json data failed,", err)
		return
	}

	data := packSendData(packetData)
	fmt.Println("Send packet:", string(packetData))
	fmt.Println("Byte data:", data)

	if _, err := c.conn.Write(data); err != nil {
		log.Fatalln("Send request failed,", err)
		return
	}
}

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
