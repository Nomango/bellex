package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Schedule))
}

type Schedule struct {
	Id      int
	Content string `orm:"size(130)"`

	User     *User      `orm:"rel(fk)"`
	Mechines []*Mechine `orm:"reverse(many)" json:"-"`

	CreateTime time.Time `orm:"auto_now_add"`
	UpdateTime time.Time `orm:"auto_now"`
}

// Insert ...
func (b *Schedule) Insert() error {
	_, err := orm.NewOrm().Insert(b)
	return err
}

// Insert ...
func (b *Schedule) InsertOrUpdate() error {
	_, err := orm.NewOrm().InsertOrUpdate(b)
	return err
}

// Read ...
func (b *Schedule) Read(fields ...string) error {
	return orm.NewOrm().Read(b, fields...)
}

// Update ...
func (b *Schedule) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(b, fields...)
	return err
}

// Delete ...
func (b *Schedule) Delete() error {
	_, err := orm.NewOrm().Delete(b)
	return err
}

func Schedules() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*Schedule)(nil)).OrderBy("-CreateTime")
}
