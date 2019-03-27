// Copyright (C) 2018 Nomango - All Rights Reserved

package tcppacket

import (
	"errors"
	"log"
	"regexp"

	"github.com/nomango/bellex/server/services/tcp/types"
)

const (
	baseRegStr = `id:([\w]+);code:([\w]+);req:([\w]+);(.*)`
)

var (
	baseRegExp *regexp.Regexp

	packetTypes = map[string]byte{
		"connect":      types.PacketTypeConnect,
		"request_time": types.PacketTypeRequestTime,
		"heart_beat":   types.PacketTypeHeartBeat,
	}
)

// LoadPacket parses Packet-encoding data
func LoadPacket(req string) (*types.Packet, error) {
	if matched := baseRegExp.MatchString(req); !matched {
		log.Println("Invalid request", req)
		return nil, errors.New("Invalid request")
	}

	params := baseRegExp.FindStringSubmatch(req)
	if len(params) != 5 {
		return nil, errors.New("Invalid request")
	}

	packetType, ok := packetTypes[params[3]]
	if !ok {
		log.Println("Unknown request type", req)
		return nil, errors.New("Unknown request type")
	}

	packet := &types.Packet{
		Auth: types.AuthPacket{
			Code:   params[1],
			Secret: params[2],
		},
		Type: packetType,
	}
	return packet, nil
}

func init() {
	baseRegExp = regexp.MustCompile(baseRegStr)
}
