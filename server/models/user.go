package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/nomango/bellex/server/modules/utils"
)

func init() {
	orm.RegisterModel(new(User))
}

const (
	UserRoleNormal = iota
	UserRoleAdmin
)

type User struct {
	Id       int    `json:"id"`
	UserName string `orm:"size(30);unique" json:"username"`
	Password string `orm:"size(128)" json:"-"`
	Email    string `orm:"size(80);unique" json:"email"`
	Role     int    `orm:"index;default(0)" json:"role"`
	IsForbid bool   `orm:"index" json:"is_forbid"`

	Institution *Institution `orm:"rel(fk)" json:"institution"`

	CreateTime time.Time `orm:"auto_now_add" json:"create_time"`
	UpdateTime time.Time `orm:"auto_now" json:"update_time"`
}

func (u *User) IsNormal() bool {
	return u.Role == UserRoleNormal
}

func (u *User) IsAdmin() bool {
	return u.Role == UserRoleAdmin
}

// Insert ...
func (u *User) Insert() error {
	_, err := orm.NewOrm().Insert(u)
	return err
}

// Insert ...
func (u *User) InsertOrUpdate() error {
	_, err := orm.NewOrm().InsertOrUpdate(u)
	return err
}

// Read ...
func (u *User) Read(fields ...string) error {
	return orm.NewOrm().Read(u, fields...)
}

// Update ...
func (u *User) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(u, fields...)
	return err
}

// Delete ...
func (u *User) Delete() error {
	_, err := orm.NewOrm().Delete(u)
	return err
}

// SetNewPassword ...
func (u *User) SetNewPassword(password string) {
	if len(password) == 0 {
		return
	}
	salt := GetUserSalt()
	u.Password = fmt.Sprintf("%s$%s", salt, utils.EncodePassword(password, salt))
}

// SaveNewPassword ...
func (u *User) SaveNewPassword(password string) error {
	u.SetNewPassword(password)
	return u.Update("Password", "UpdateTime")
}

// Users returns all users query
func Users() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*User)(nil))
}

// GetUserSalt returns a user salt token
func GetUserSalt() string {
	return utils.RandString(10)
}

// HasUser returns whether if user exists
func HasUser(usernameOrEmail string) bool {
	users := Users()
	if strings.IndexRune(usernameOrEmail, '@') == -1 {
		return users.Filter("UserName", usernameOrEmail).Exist()
	}
	return users.Filter("Email", usernameOrEmail).Exist()
}

// FindUser ...
func FindUser(usernameOrEmail string) (*User, error) {
	var user User
	users := Users()

	if strings.IndexRune(usernameOrEmail, '@') == -1 {
		err := users.Filter("UserName", usernameOrEmail).One(&user)
		return &user, err
	}

	err := users.Filter("Email", usernameOrEmail).One(&user)
	return &user, err
}

// RegisterUser ...
func RegisterUser(user *User, username, email, password string) error {
	user.UserName = strings.ToLower(username)
	user.Email = strings.ToLower(email)
	user.SetNewPassword(password)

	user.IsForbid = false
	user.Role = UserRoleNormal

	return user.Insert()
}

// CreateAdminUser register an administrator
func CreateAdminUser() error {
	user := &User{UserName: "admin"}
	var err error
	if !HasUser(user.UserName) {
		if err = RegisterUser(user, user.UserName, "admin@localhost.com", "admin"); err == nil {
			user.Role = UserRoleAdmin
			user.IsForbid = false
			user.Institution = &Institution{Id: 1}
			err = user.Update("Role", "IsForbid", "Institution")
		}
	}
	return err
}
