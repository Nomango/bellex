// Copyright (C) 2018 Nomango - All Rights Reserved

package types

import "github.com/nomango/bellex/server/models"

// Packet types
const (
	PacketTypeConnect byte = iota
	PacketTypeRequestTime
	PacketTypeHeartBeat
)

// AuthPacket contains authority verification infofmation
type AuthPacket struct {
	Code   string `json:"id"`
	Secret string `json:"code"`
}

// Packet send & recive data format
type Packet struct {
	Auth AuthPacket `json:"auth"`
	Type byte       `json:"type"`
	Data string     `json:"data"`
}

// DefaultPacket returns a default packet
func DefaultPacket() *Packet {
	return &Packet{
		Auth: AuthPacket{
			Code:   "BW123456",
			Secret: "Xwa8pj7z",
		},
		Type: 0,
		Data: "",
	}
}

// GetMechine returns related mechine
func (p *Packet) GetMechine() (*models.Mechine, error) {
	var mechine models.Mechine
	if err := models.Mechines().Filter("Code", p.Auth.Code).One(&mechine); err != nil {
		return nil, err
	}
	return &mechine, nil
}
