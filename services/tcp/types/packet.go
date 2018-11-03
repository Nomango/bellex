// Copyright (C) 2018 Nomango - All Rights Reserved

package types

// Packet send & recive data format
type Packet struct {
	Type byte
	Data []byte
}

// Packet types
const (
	PacketTypeGetTime byte = iota
	PacketTypeChangeMode
)

var (
	// PacketTypes all types of packet
	PacketTypes = [...]byte{
		PacketTypeGetTime,
		PacketTypeChangeMode,
	}
)

// VerifyPacket contains authority verification infomation
type VerifyPacket struct {
	BellID     string `json:"bell_id"`
	UniqueCode string `json:"unique_code"`
}

// TimePacket contains time info
type TimePacket struct {
	Timestamp int64 `json:"timestamp"`
}
