package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/nomango/bellex/server/modules/utils"
)

func init() {
	orm.RegisterModel(new(Mechine))
}

type Mechine struct {
	Id      int                `json:"id"`
	Name    string             `orm:"size(30)" json:"name"`
	Code    string             `orm:"unique;size(8)" json:"code"`
	Secret  string             `orm:"size(8)" json:"secret"`
	Idle    bool               `json:"idle"`
	Accept  bool               `orm:"-" json:"accept"`
	Connect *MechineConnection `orm:"-" json:"-"`

	Institution *Institution `orm:"rel(fk)" json:"institution"`
	Schedule    *Schedule    `orm:"rel(fk)" json:"schedule"`

	CreateTime time.Time `orm:"auto_now_add" json:"create_time"`
	UpdateTime time.Time `orm:"auto_now" json:"update_time"`
}

// Insert ...
func (m *Mechine) Insert() error {
	_, err := orm.NewOrm().Insert(m)
	return err
}

// InsertOrUpdate ...
func (m *Mechine) InsertOrUpdate() error {
	_, err := orm.NewOrm().InsertOrUpdate(m)
	return err
}

// Read ...
func (m *Mechine) Read(fields ...string) error {
	return orm.NewOrm().Read(m, fields...)
}

// Update ...
func (m *Mechine) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(m, fields...)
	return err
}

// Delete ...
func (m *Mechine) Delete() error {
	_, err := orm.NewOrm().Delete(m)
	return err
}

// UpdateStatus ...
func (m *Mechine) UpdateStatus() {
	if conn, err := GetConnection(m); err != nil {
		m.Connect = conn
		m.Accept = false
	} else {
		m.Connect = nil
		m.Accept = true
	}
}

// SetNewSecret ...
func (m *Mechine) SetNewSecret() {
	m.Secret = utils.RandString(8)
}

// SaveNewSecret ...
func (m *Mechine) SaveNewSecret() error {
	m.SetNewSecret()
	return m.Update("Secret")
}

func Mechines() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*Mechine)(nil)).OrderBy("-CreateTime")
}

func MechinesAccepted() orm.QuerySeter {
	return Mechines().Filter("Accept", true)
}

func MechinesUnaccepted() orm.QuerySeter {
	return Mechines().Filter("Accept", false)
}
