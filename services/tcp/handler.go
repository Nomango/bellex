// Copyright (C) 2018 Nomango - All Rights Reserved

package tcp

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/nomango/bellex/services/ntp"
	tcpPacket "github.com/nomango/bellex/services/tcp/packet"
	"github.com/nomango/bellex/services/tcp/types"
)

// PacketHandler handle request packets
func PacketHandler(req []byte, conn net.Conn, outputCh chan<- []byte, closeCh chan<- struct{}) {

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
		log.Println("Permission denied:", string(req))
		err = errors.New("Permission denied")
		return
	} else {
		log.Println("Receive request:", string(req))
	}

	switch packet.Type {
	case types.PacketTypeConnect:
		response, err = handleRequestConnect(packet, conn, outputCh, closeCh)
	case types.PacketTypeRequestTime:
		response, err = handleRequestTime()
	case types.PacketTypeHeartBeat:
		response, err = handleRequestHeartBeat(packet)
	case types.PacketTypeGetSchedule:
		response, err = handleRequestGetSchedule(packet)
	default:
		err = errors.New("Invalid request")
	}
}

func handleRequestConnect(packet *types.Packet, conn net.Conn, outputCh chan<- []byte, closeCh chan<- struct{}) (string, error) {

	mechine, err := packet.GetMechine()
	if err != nil {
		return "", errors.New("Permission denied")
	}

	if types.ExistsConnection(mechine.Code) {
		return "", errors.New("Connection already exists")
	}

	// connection already exists
	mechine.SaveNewSecret()

	if err := types.AddConnection(mechine.Code, conn, outputCh, closeCh); err != nil {
		log.Println("Add connection failed", err)
	}

	return "unique_code:" + mechine.Secret + ";", nil
}

func handleRequestTime() (string, error) {
	now, err := sendNTPRequest()
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

func handleRequestHeartBeat(packet *types.Packet) (string, error) {
	mechine, err := packet.GetMechine()
	if err != nil {
		return "", errors.New("Permission denied")
	}

	if !types.ExistsConnection(mechine.Code) {
		return "", errors.New("Connection not found")
	}

	switch {
	case strings.Contains(packet.Data, "status:idle"):
		mechine.Idle = true
		mechine.Update("Idle")
	case strings.Contains(packet.Data, "status:working"):
		mechine.Idle = false
		mechine.Update("Idle")
	default:
		return "status:0;", nil
	}

	return "status:1;", nil
}

func handleRequestGetSchedule(packet *types.Packet) (string, error) {
	mechine, err := packet.GetMechine()
	if err != nil {
		return "", errors.New("Permission denied")
	}

	if !types.ExistsConnection(mechine.Code) {
		return "", errors.New("Connection not found")
	}

	return "schedule:" + mechine.Schedule.FormatContent() + ";", nil
}

func sendNTPRequest() (time.Time, error) {
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
