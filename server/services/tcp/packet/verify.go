// Copyright (C) 2018 Nomango - All Rights Reserved

package tcppacket

import (
	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/services/tcp/types"
)

// Verify check if the client has permissions
func Verify(packet *types.Packet) bool {
	if packet.Type == types.PacketTypeConnect {
		return true // ignore connect verify
	}
	bells := models.GetAllBells()
	if bell, ok := bells[packet.Auth.ID]; ok {
		return bell.ID == packet.Auth.ID && bell.Code == packet.Auth.Code
	}
	return false
}
