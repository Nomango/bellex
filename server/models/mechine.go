package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Mechine))
}

type Mechine struct {
	Id     int
	Code   string `orm:"unique;size(30)"`
	Secret string `orm:"size(30)"`

	User     *User     `orm:"rel(fk)"`
	Schedule *Schedule `orm:"rel(fk)"`

	CreateTime time.Time `orm:"auto_now_add"`
	UpdateTime time.Time `orm:"auto_now"`
}

// NewMechine ...
func NewMechine(Code string, Secret string) *Mechine {
	return &Mechine{
		Code:   Code,
		Secret: Secret,
	}
}

// Insert ...
func (b *Mechine) Insert() error {
	_, err := orm.NewOrm().Insert(b)
	return err
}

// Insert ...
func (b *Mechine) InsertOrUpdate() error {
	_, err := orm.NewOrm().InsertOrUpdate(b)
	return err
}

// Read ...
func (b *Mechine) Read(fields ...string) error {
	return orm.NewOrm().Read(b, fields...)
}

// Update ...
func (b *Mechine) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(b, fields...)
	return err
}

// Delete ...
func (b *Mechine) Delete() error {
	_, err := orm.NewOrm().Delete(b)
	return err
}

func Mechines() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*Mechine)(nil)).OrderBy("-CreateTime")
}
