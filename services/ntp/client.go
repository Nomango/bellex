// Copyright (C) 2018 Nomango - All Rights Reserved

package ntp

import (
	"encoding/binary"
	"log"
	"net"
	"time"
)

// NTP Servers
const (
	ServerAliyun  = "ntp1.aliyun.com:123"
	ServerNTSC    = "ntp.ntsc.ac.cn:123"
	ServerWindows = "time.windows.com:123"
)

var (
	// Servers defines all NTP servers
	Servers = [...]string{
		ServerAliyun,
		ServerNTSC,
		ServerWindows,
	}
)

// Client a client connect to NTP service
type Client struct {
	conn net.Conn
}

// NewClient returns a new NTP client
func NewClient(host string) (*Client, error) {
	conn, err := net.Dial("udp", host)
	if err != nil {
		return nil, err
	}
	if err := conn.SetDeadline(time.Now().Add(2 * time.Second)); err != nil {
		return nil, err
	}
	return &Client{
		conn: conn,
	}, nil
}

// Close stop connection
func (c *Client) Close() {
	c.conn.Close()
}

// Request send request and get response
func (c *Client) Request() (*Packet, error) {
	req := DefaultPacket()
	if err := binary.Write(c.conn, binary.BigEndian, req); err != nil {
		return nil, err
	}

	rsp := &Packet{}
	if err := binary.Read(c.conn, binary.BigEndian, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

// SendRequest ...
func SendRequest(host string) (time.Time, error) {
	client, err := NewClient(host)
	if err != nil {
		log.Println("Start NTP client failed:", err)
		return time.Time{}, err
	}

	defer client.Close()

	packet, err := client.Request()
	if err != nil {
		log.Println("Failed to send request:", err)
		return time.Time{}, err
	}

	return packet.Parse(), nil
}
