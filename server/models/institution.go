package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Institution))
}

type Institution struct {
	Id   int    `json:"id"`
	Name string `orm:"unique;size(30)" json:"name"`

	Users     []*User     `orm:"reverse(many)" json:"-"`
	Mechines  []*Mechine  `orm:"reverse(many)" json:"-"`
	Schedules []*Schedule `orm:"reverse(many)" json:"-"`

	CreateTime time.Time `orm:"auto_now_add" json:"create_time"`
	UpdateTime time.Time `orm:"auto_now" json:"update_time"`
}

// Insert ...
func (i *Institution) Insert() error {
	_, err := orm.NewOrm().Insert(i)
	return err
}

// Insert ...
func (i *Institution) InsertOrUpdate() error {
	_, err := orm.NewOrm().InsertOrUpdate(i)
	return err
}

// Read ...
func (i *Institution) Read(fields ...string) error {
	return orm.NewOrm().Read(i, fields...)
}

// Update ...
func (i *Institution) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(i, fields...)
	return err
}

// Delete ...
func (i *Institution) Delete() error {
	_, err := orm.NewOrm().Delete(i)
	return err
}

func Institutions() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*Institution)(nil)).OrderBy("-CreateTime")
}
