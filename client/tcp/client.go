// Copyright (C) 2018 Nomango - All Rights Reserved

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Client tcp client
type Client struct {
	conn *net.TCPConn
	addr *net.TCPAddr
	stop chan struct{}
}

// NewClient returns a new tcp cliet
func NewClient(ip string) (*Client, error) {
	addr, err := net.ResolveTCPAddr("tcp", ip)
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

// Receiver returns a Reader to get server response
func (c *Client) Receiver() *bufio.Reader {
	return bufio.NewReader(c.conn)
}

// RequestTime send 'GetTime' request
func (c *Client) RequestTime() {

	packetData := []byte(`id:123;code:123;req:schedule;`)
	//packetData := []byte(`id:123;code:123;req:request_timing;`)
	fmt.Println("Send packet:", string(packetData))

	data := packSendData(packetData)
	if _, err := c.conn.Write(data); err != nil {
		log.Fatalln("Send request failed,", err)
		return
	}
}

func packSendData(sendBytes []byte) []byte {
	size := len(sendBytes) + 5
	result := make([]byte, size)
	result[0] = 0xFF
	result[1] = 0xFF
	result[2] = byte(uint16(len(sendBytes)) & 0xFF)

	copy(result[3:], sendBytes)

	result[size-2] = 0xFF
	result[size-1] = 0xFE
	return result
}
