// Copyright (C) 2018 Nomango - All Rights Reserved

package tcp

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/nomango/bellex/services/ntp"
	tcpPacket "github.com/nomango/bellex/services/tcp/packet"
	"github.com/nomango/bellex/services/tcp/types"
)

// PacketHandler handle request packets
func PacketHandler(req []byte, conn net.Conn) {

	packet, err := tcpPacket.LoadPacket(string(req))
	if err != nil {
		write("Invalid request", conn)
		return
	}

	if !tcpPacket.Verify(packet) {
		write("Permission denied", conn)
		return
	}

	switch packet.Type {
	case types.PacketTypeRequestTime:
		if now, err := requestNTP(); err != nil {
			write(err.Error(), conn)
		} else {
			week := int(now.Weekday())
			if week == 0 {
				week = 7
			}
			response := fmt.Sprintf("current_time:%s%02d%s", now.Format("0504150201"), week, now.Format("06"))
			write(response, conn)
		}
	case types.PacketTypeSchedule:
		write("schedule:1234567890123456", conn)
	default:
		write("Invalid request", conn)
	}
}

func requestNTP() (time.Time, error) {
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

	var result time.Time
	for range make([]int, size) {
		result = <-signals
		if !result.IsZero() {
			break
		}
	}

	if result.IsZero() {
		return time.Time{}, errors.New("Request NTP failed")
	}
	return result, nil
}

func write(data string, conn net.Conn) error {
	bytes := append([]byte(data), byte(0))
	if _, err := conn.Write(bytes); err != nil {
		return err
	}
	return nil
}
