package tcprpc

import (
	"github.com/nomango/bellex/services/tcp/types"
)

type Connection struct {
}

type ConnectionCommonReq struct {
	Code string
}

type ConnectionSendDataReq struct {
	Code string
	Data string
}

func (c *Connection) Exists(req ConnectionCommonReq, res *bool) error {
	*res = types.ExistsConnection(req.Code)
	return nil
}

func (c *Connection) SendData(req ConnectionSendDataReq, res *struct{}) error {
	conn, err := types.GetConnection(req.Code)
	if err != nil {
		return err
	}
	if err := conn.SendData(req.Data); err != nil {
		return err
	}
	return nil
}

func (c *Connection) Close(req ConnectionCommonReq, res *struct{}) error {
	conn, err := types.GetConnection(req.Code)
	if err != nil {
		return err
	}
	if err := conn.CloseConnection(); err != nil {
		return err
	}
	return nil
}
