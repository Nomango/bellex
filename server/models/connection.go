package models

import (
	"errors"
	"net"
	"sync"
)

var (
	connects map[int]*BellConnection
	mutex    sync.Mutex
)

type BellConnection struct {
	Id   int
	Code string
	Conn net.Conn
}

func AddConnection(bell *Bell, conn net.Conn) error {
	mutex.Lock()
	defer mutex.Unlock()

	if ExistsConnection(bell.Id) {
		return errors.New("Connection already exists")
	}

	connects[bell.Id] = &BellConnection{
		Id:   bell.Id,
		Code: bell.Code,
		Conn: conn,
	}
	return nil
}

func GetConnection(bell *Bell) (net.Conn, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if !ExistsConnection(bell.Id) {
		return nil, errors.New("Connection not exists")
	}

	return connects[bell.Id].Conn, nil
}

func DeleteConnection(bell *Bell) error {
	mutex.Lock()
	defer mutex.Unlock()

	if !ExistsConnection(bell.Id) {
		return errors.New("Connection not exists")
	}

	delete(connects, bell.Id)
	return nil
}

func ExistsConnection(Id int) bool {
	_, ok := connects[Id]
	return ok
}

func init() {
	connects = make(map[int]*BellConnection)
}
