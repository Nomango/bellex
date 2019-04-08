package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Insititution))
}

type Insititution struct {
	Id   int    `json:"id"`
	Name string `orm:"unique;size(30)" json:"name"`

	Users []*User `orm:"reverse(many)" json:"-"`
	Mechines  []*Mechine  `orm:"reverse(many)" json:"-"`
	Schedules []*Schedule `orm:"reverse(many)" json:"-"`

	CreateTime time.Time `orm:"auto_now_add" json:"-"`
	UpdateTime time.Time `orm:"auto_now" json:"-"`
}

// Insert ...
func (i *Insititution) Insert() error {
	_, err := orm.NewOrm().Insert(i)
	return err
}

// Insert ...
func (i *Insititution) InsertOrUpdate() error {
	_, err := orm.NewOrm().InsertOrUpdate(i)
	return err
}

// Read ...
func (i *Insititution) Read(fields ...string) error {
	return orm.NewOrm().Read(i, fields...)
}

// Update ...
func (i *Insititution) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(i, fields...)
	return err
}

// Delete ...
func (i *Insititution) Delete() error {
	_, err := orm.NewOrm().Delete(i)
	return err
}

func Insititutions() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*Insititution)(nil)).OrderBy("-CreateTime")
}
