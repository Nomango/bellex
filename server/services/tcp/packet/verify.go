// Copyright (C) 2018 Nomango - All Rights Reserved

package tcppacket

import (
	"errors"

	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/services/tcp/types"
)

// Verify check if the client has permissions
func Verify(packet *types.Packet) error {
	if packet.Type == types.PacketTypeConnect {
		return nil // ignore connect
	}
	bell := &models.Bell{Code: packet.Auth.Code}
	if err := bell.Read("Code"); err != nil {
		return err
	}
	if bell.Code == packet.Auth.Code && bell.Secret == packet.Auth.Secret {
		return nil
	}
	return errors.New("Verify secret failed")
}
