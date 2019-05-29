package tcprpc

import (
	"net"
	"net/rpc"

	"github.com/nomango/bellex/server/modules/settings"
)

// NewServer returns a new rpc server
func NewServer() (lis net.Listener, err error) {
	rpcServer := rpc.NewServer()
	if err := rpcServer.Register(new(Connection)); err != nil {
		return nil, err
	}

	rpcServer.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	if lis, err := net.Listen("tcp", settings.LocalAddr+":"+settings.RpcPort); err != nil {
		return nil, err
	} else {
		return lis, nil
	}
}
