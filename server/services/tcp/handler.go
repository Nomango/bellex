// Copyright (C) 2018 Nomango - All Rights Reserved

package tcp

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/nomango/bellex/server/services/ntp"
	tcpPacket "github.com/nomango/bellex/server/services/tcp/packet"
	"github.com/nomango/bellex/server/services/tcp/types"
	"github.com/nomango/bellex/server/services/utils"
)

// PacketHandler handle request packets
func PacketHandler(req []byte, conn net.Conn) {

	var (
		packet   *types.Packet
		response string
		err      error
	)

	defer func() {
		if err != nil {
			write(err.Error(), conn)
		} else {
			write(response, conn)
		}
	}()

	packet, err = tcpPacket.LoadPacket(string(req))
	if err != nil {
		err = errors.New("error:Invalid request;#")
		return
	}

	if !tcpPacket.Verify(packet) {
		err = errors.New("error:Permission denied;#")
		return
	}

	switch packet.Type {
	case types.PacketTypeConnect:
		response, err = requestConnect()
	case types.PacketTypeRequestTime:
		response, err = requestTime()
	case types.PacketTypeHeartBeat:
		response = "status:1;#"
	default:
		err = errors.New("error:Invalid request;#")
	}
}

func requestConnect() (string, error) {
	code := utils.RandString(8)
	// Check whether code exists
	// Save code
	return "unique_code:" + code + ";#", nil
}

func requestTime() (string, error) {
	now, err := requestNTP()
	if err != nil {
		return "", err
	}

	week := int(now.Weekday())
	if week == 0 {
		week = 7
	}
	response := fmt.Sprintf("current_time:%s%02d%s;#", now.Format("0504150201"), week, now.Format("06"))
	return response, nil
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
		return time.Time{}, errors.New("error:Request NTP failed;#")
	}
	return result, nil
}

func write(data string, conn net.Conn) error {
	if _, err := conn.Write([]byte(data)); err != nil {
		return err
	}
	return nil
}
