package models

import (
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Schedule))
}

type Schedule struct {
	Id      int    `json:"id"`
	Name    string `orm:"size(30)" json:"name"`
	Content string `orm:"size(130)" json:"content"`

	Institution *Institution `orm:"rel(fk)" json:"institution"`
	Mechines    []*Mechine   `orm:"reverse(many)" json:"-"`

	CreateTime time.Time `orm:"auto_now_add" json:"create_time"`
	UpdateTime time.Time `orm:"auto_now" json:"update_time"`
}

// Insert ...
func (s *Schedule) Insert() error {
	_, err := orm.NewOrm().Insert(s)
	return err
}

// Insert ...
func (s *Schedule) InsertOrUpdate() error {
	_, err := orm.NewOrm().InsertOrUpdate(s)
	return err
}

// Read ...
func (s *Schedule) Read(fields ...string) error {
	return orm.NewOrm().Read(s, fields...)
}

// Update ...
func (s *Schedule) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(s, fields...)
	return err
}

// Delete ...
func (s *Schedule) Delete() error {
	_, err := orm.NewOrm().Delete(s)
	return err
}

// FormatContent ...
func (s *Schedule) FormatContent() string {
	var timetable string
	for _, time := range strings.Split(s.Content, " ") {
		for _, num := range strings.Split(time, ":") {
			timetable += num
		}
	}
	return timetable
}

func Schedules() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*Schedule)(nil)).OrderBy("-CreateTime")
}
