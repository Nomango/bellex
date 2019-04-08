// Copyright (C) 2018 Nomango - All Rights Reserved

package tcp

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/astaxie/beego"

	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/utils"
	"github.com/nomango/bellex/server/services/ntp"
	tcpPacket "github.com/nomango/bellex/server/services/tcp/packet"
	"github.com/nomango/bellex/server/services/tcp/types"
)

// PacketHandler handle request packets
func PacketHandler(req []byte, conn net.Conn, outputCh chan<- []byte) {

	var (
		packet   *types.Packet
		response string
		err      error
	)

	defer func() {
		if err != nil {
			outputCh <- []byte("error:" + err.Error() + ";")
		} else {
			outputCh <- []byte(response)
		}
	}()

	packet, err = tcpPacket.LoadPacket(string(req))
	if err != nil {
		return
	}

	err = tcpPacket.Verify(packet)
	if err != nil {
		beego.Error("Permission denied", string(req))
		err = errors.New("Permission denied")
		return
	}

	switch packet.Type {
	case types.PacketTypeConnect:
		response, err = requestConnect(packet, conn)
	case types.PacketTypeRequestTime:
		response, err = requestTime()
	case types.PacketTypeHeartBeat:
		response = "status:1;"
	default:
		err = errors.New("Invalid request")
	}
}

func requestConnect(packet *types.Packet, conn net.Conn) (string, error) {
	// if _, ok := bells[packet.Auth.Code]; ok {
	// 	return tcpPacket.NewError("Code exists")
	// }
	code := utils.RandString(8)
	mechine := models.NewMechine(packet.Auth.Code, code)
	if err := mechine.Insert(); err != nil {
		return "", err
	}
	if err := models.AddConnection(mechine, conn); err != nil {
		beego.Error("Add connection failed", err)
	}
	return "unique_code:" + code + ";", nil
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
	response := fmt.Sprintf("current_time:%s%02d%s;", now.Format("0504150201"), week, now.Format("06"))
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
		return time.Time{}, errors.New("Request NTP failed")
	}
	return result, nil
}
