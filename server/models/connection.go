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

type MechineConnection struct {
	Id   int
	Code string
	Conn net.Conn
}

func AddConnection(mechine *Mechine, conn net.Conn) error {
	mutex.Lock()
	defer mutex.Unlock()

	if ExistsConnection(mechine.Id) {
		return errors.New("Connection already exists")
	}

	connects[mechine.Id] = &MechineConnection{
		Id:   mechine.Id,
		Code: mechine.Code,
		Conn: conn,
	}
	return nil
}

func GetConnection(mechine *Mechine) (net.Conn, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if !ExistsConnection(mechine.Id) {
		return nil, errors.New("Connection not exists")
	}

	return connects[mechine.Id].Conn, nil
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

func ExistsConnection(Id int) bool {
	_, ok := connects[Id]
	return ok
}

func init() {
	connects = make(map[int]*MechineConnection)
}
