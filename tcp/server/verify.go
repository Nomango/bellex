// Copyright (C) 2018 Nomango - All Rights Reserved

package server

import (
	"log"

	"github.com/nomango/bellex/services/tcp/types"
)

// Verify check if the client has permissions
func Verify(packet *types.Packet) bool {
	// FIX ME!!! @Nomango
	log.Println(*packet)
	return packet.Auth.ID == "123" && packet.Auth.Code == "123"
}
