package tcprpc

import (
	"errors"
	"net"
	"net/rpc"

	"github.com/nomango/bellex/server/modules/settings"
)

// getLocalIP get local ip address
func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("Local IP not found")
}

// NewServer returns a new rpc server
func NewServer() (lis net.Listener, err error) {
	rpcServer := rpc.NewServer()
	if err := rpcServer.Register(new(Connection)); err != nil {
		return nil, err
	}

	rpcServer.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	var addr string
	if settings.Debug {
		addr = "127.0.0.1:" + settings.RpcPort
	} else {
		var localIP string
		localIP, err = getLocalIP()
		if err != nil {
			return nil, err
		}
		addr = localIP + ":" + settings.RpcPort
	}

	if lis, err := net.Listen("tcp", addr); err != nil {
		return nil, err
	} else {
		return lis, nil
	}
}
