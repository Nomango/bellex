package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Bell struct {
	Id      int
	Code    string    `orm:"unique;size(30)"`
	Secret  string    `orm:"size(30)"`
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
}

// NewBell ...
func NewBell(Code string, Secret string) *Bell {
	return &Bell{
		Code:   Code,
		Secret: Secret,
	}
}

// Insert ...
func (b *Bell) Insert() error {
	_, err := orm.NewOrm().Insert(b)
	return err
}

// Read ...
func (b *Bell) Read(fields ...string) error {
	return orm.NewOrm().Read(b, fields...)
}

// Update ...
func (b *Bell) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(b, fields...)
	return err
}

// Delete ...
func (b *Bell) Delete() error {
	_, err := orm.NewOrm().Delete(b)
	return err
}

func Bells() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*Bell)(nil)).OrderBy("-Created")
}
