package types

import (
	"errors"
	"net"
	"sync"
)

var (
	connects map[string]*MechineConnection
	mutex    sync.RWMutex
)

func init() {
	connects = make(map[string]*MechineConnection)
}

type MechineConnection struct {
	Code     string
	Conn     net.Conn
	OutputCh chan<- []byte
	CloseCh  chan<- struct{}
}

func (c *MechineConnection) SendData(data string) error {
	if c.OutputCh == nil {
		return errors.New("no connection")
	}

	c.OutputCh <- []byte(data)
	return nil
}

func (c *MechineConnection) CloseConnection() error {
	if c.CloseCh == nil {
		return errors.New("no connection")
	}

	c.CloseCh <- struct{}{}
	return nil
}

func AddConnection(code string, conn net.Conn, output chan<- []byte, close chan<- struct{}) error {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := connects[code]; ok {
		return errors.New("Connection already exists")
	}

	connects[code] = &MechineConnection{
		Code:     code,
		Conn:     conn,
		OutputCh: output,
		CloseCh:  close,
	}
	return nil
}

func GetConnection(code string) (*MechineConnection, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	if _, ok := connects[code]; !ok {
		return nil, errors.New("Connection not exists")
	}

	return connects[code], nil
}

func DeleteConnection(code string) error {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := connects[code]; !ok {
		return errors.New("Connection not exists")
	}

	// mechine.CloseConnection()

	delete(connects, code)
	return nil
}

func DeleteConnectionWithConn(conn net.Conn) error {
	mutex.Lock()
	defer mutex.Unlock()

	for key, c := range connects {
		if c.Conn.RemoteAddr().String() == conn.RemoteAddr().String() {
			delete(connects, key)
			return nil
		}
	}

	return errors.New("Connection not exists")
}

func ExistsConnection(code string) bool {
	mutex.RLock()
	defer mutex.RUnlock()

	_, ok := connects[code]
	return ok
}
