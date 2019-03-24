package models

import (
	"errors"
	"net"
	"strconv"
	"sync"
	"time"
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

func AddBell(bell Bell) (ID string) {
	bell.ID = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Bells[bell.ID] = &bell
	return bell.ID
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
