// Copyright (C) 2018 Nomango - All Rights Reserved

package tcppacket

import (
	"fmt"

	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/services/tcp/types"
)

// Verify check if the client has permissions
func Verify(packet *types.Packet) (bool, error) {
	if packet.Type == types.PacketTypeConnect {
		return true, nil // ignore connect verify
	}
	bells := models.GetAllBells()
	if bell, ok := bells[packet.Auth.ID]; ok {
		fmt.Println(bell.ID, bell.Code)
		return bell.ID == packet.Auth.ID && bell.Code == packet.Auth.Code, nil
	}
	return false, nil
}
