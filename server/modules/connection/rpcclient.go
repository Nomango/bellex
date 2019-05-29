package connection

import (
	"net/rpc"

	"github.com/nomango/bellex/server/modules/settings"
)

const (
	ConnGet      = "Connection.Get"
	ConnExists   = "Connection.Exists"
	ConnSendData = "Connection.SendData"
	ConnClose    = "Connection.Close"
)

type ConnectionCommonReq struct {
	Code string
}

type ConnectionSendDataReq struct {
	Code string
	Data string
}

// NewClient returns a new rpc client
func NewClient() (*rpc.Client, error) {
	if cli, err := rpc.DialHTTP("tcp", settings.RpcAddr); err != nil {
		return nil, err
	} else {
		return cli, nil
	}
}

func Exists(code string) (bool, error) {
	cli, err := NewClient()
	if err != nil {
		return false, err
	}
	defer cli.Close()

	var exists bool
	if err := cli.Call(ConnExists, ConnectionCommonReq{Code: code}, &exists); err != nil {
		return false, err
	}
	return exists, nil
}

func SendData(code string, data string) error {
	cli, err := NewClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	if err := cli.Call(ConnSendData, ConnectionSendDataReq{Code: code, Data: data}, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func Close(code string) error {
	cli, err := NewClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	if err := cli.Call(ConnClose, ConnectionCommonReq{Code: code}, &struct{}{}); err != nil {
		return err
	}
	return nil
}
