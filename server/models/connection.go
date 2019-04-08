package models

import (
	"errors"
	"net"
	"sync"
)

var (
	connects map[int]*MechineConnection
	mutex    sync.Mutex
)

func init() {
	connects = make(map[int]*MechineConnection)
}

type MechineConnection struct {
	Id     int
	Code   string
	Conn   net.Conn
	Output chan<- []byte
}

func AddConnection(mechine *Mechine, conn net.Conn, output chan<- []byte) error {
	mutex.Lock()
	defer mutex.Unlock()

	if ExistsConnection(mechine.Id) {
		return errors.New("Connection already exists")
	}

	connects[mechine.Id] = &MechineConnection{
		Id:     mechine.Id,
		Code:   mechine.Code,
		Conn:   conn,
		Output: output,
	}
	return nil
}

func GetConnection(mechine *Mechine) (*MechineConnection, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if !ExistsConnection(mechine.Id) {
		return nil, errors.New("Connection not exists")
	}

	return connects[mechine.Id], nil
}

func DeleteConnection(mechine *Mechine) error {
	mutex.Lock()
	defer mutex.Unlock()

	if !ExistsConnection(mechine.Id) {
		return errors.New("Connection not exists")
	}

	delete(connects, mechine.Id)
	return nil
}

func DeleteConnectionWithConn(conn net.Conn) error {
	mutex.Lock()
	defer mutex.Unlock()

	for key, c := range connects {
		if c.Conn == conn {
			delete(connects, key)
			return nil
		}
	}

	return errors.New("Connection not exists")
}

func ExistsConnection(Id int) bool {
	_, ok := connects[Id]
	return ok
}
