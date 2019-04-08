package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Mechine))
}

type Mechine struct {
	Id     int    `json:"id"`
	Code   string `orm:"unique;size(30)" json:"code"`
	Secret string `orm:"size(30)" json:"secret"`

	Insititution *Insititution `orm:"rel(fk)" json:"insititution"`
	Schedule     *Schedule     `orm:"rel(fk)" json:"-"`

	CreateTime time.Time `orm:"auto_now_add" json:"-"`
	UpdateTime time.Time `orm:"auto_now" json:"-"`
}

// NewMechine ...
func NewMechine(Code string, Secret string) *Mechine {
	return &Mechine{
		Code:   Code,
		Secret: Secret,
	}
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

func Mechines() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*Mechine)(nil)).OrderBy("-CreateTime")
}
