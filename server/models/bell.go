package models

import (
	"errors"
	"net"
	"sync"
)

var (
	Bells map[string]*Bell
	mutex sync.Mutex
)

type Bell struct {
	ID   string
	Code string
	conn net.Conn
}

func init() {
	Bells = make(map[string]*Bell)
}

func NewBell(ID string, Code string, conn net.Conn) *Bell {
	return &Bell{
		ID:   ID,
		Code: Code,
		conn: conn,
	}
}

func AddBell(bell *Bell) error {
	Bells[bell.ID] = bell
	return nil
}

func GetBell(ID string) (bell *Bell, err error) {
	if v, ok := Bells[ID]; ok {
		return v, nil
	}
	return nil, errors.New("ID Not Exist")
}

func GetAllBells() map[string]*Bell {
	return Bells
}

func UpdateBell(ID string, Code string) (err error) {
	if v, ok := Bells[ID]; ok {
		v.Code = Code
		return nil
	}
	return errors.New("ID Not Exist")
}

func DeleteBell(ID string) {
	delete(Bells, ID)
}
