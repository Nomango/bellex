// Copyright (C) 2018 Nomango - All Rights Reserved

package types

// Packet types
const (
	PacketTypeConnect byte = iota
	PacketTypeRequestTime
	PacketTypeHeartBeat
)

// AuthPacket contains authority verification infofmation
type AuthPacket struct {
	ID   string `json:"id"`
	Code string `json:"code"`
}

// Packet send & recive data format
type Packet struct {
	Auth AuthPacket `json:"auth"`
	Type byte       `json:"type"`
	Data []byte     `json:"data"`
}

// DefaultPacket returns a default packet
func DefaultPacket() *Packet {
	return &Packet{
		Auth: AuthPacket{
			ID:   "BW123",
			Code: "Xwa8pj7",
		},
		Type: 0,
		Data: make([]byte, 0),
	}
}
