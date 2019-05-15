package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Version))
}

type Version struct {
	Code int    `json:"code"`
	URL  string `orm:"column(url)" json:"url"`

	CreateTime time.Time `orm:"auto_now_add" json:"create_time"`
	UpdateTime time.Time `orm:"auto_now" json:"update_time"`
}

// Insert ...
func (i *Version) Insert() error {
	_, err := orm.NewOrm().Insert(i)
	return err
}

// InsertOrUpdate ...
func (i *Version) InsertOrUpdate() error {
	_, err := orm.NewOrm().InsertOrUpdate(i)
	return err
}

// Read ...
func (i *Version) Read(fields ...string) error {
	return orm.NewOrm().Read(i, fields...)
}

// Update ...
func (i *Version) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(i, fields...)
	return err
}

// Delete ...
func (i *Version) Delete() error {
	_, err := orm.NewOrm().Delete(i)
	return err
}

func GetLatestVersion() (*Version, error) {
	var ver Version
	orm.NewOrm().QueryTable((*Version)(nil)).OrderBy("-CreateTime").Limit(1).One(&ver)
	return &ver, nil
}
