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
	UserRoleSuperAdmin
)

type User struct {
	Id       int    `json:"id"`
	UserName string `orm:"size(30)" json:"username"`
	Password string `orm:"size(128)" json:"-"`
	Email    string `orm:"size(80);unique" json:"-"`
	Role     int    `orm:"index;default(0)" json:"role"`
	IsForbid bool   `orm:"index" json:"is_forbid"`
	Parent   int    `orm:"default(0)" json:"-"`

	Mechines  []*Mechine  `orm:"reverse(many)" json:"-"`
	Schedules []*Schedule `orm:"reverse(many)" json:"-"`

	CreateTime time.Time `orm:"auto_now_add" json:"-"`
	UpdateTime time.Time `orm:"auto_now" json:"-"`
}

func (u *User) IsNormal() bool {
	return u.Role == UserRoleNormal
}

func (u *User) IsAdmin() bool {
	return u.Role == UserRoleAdmin
}

func (u *User) IsSuperAdmin() bool {
	return u.Role == UserRoleSuperAdmin
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

// GetParent ...
func (u *User) GetParentAdmin() *User {
	if u.IsNormal() && u.Parent > 0 {
		parent := &User{Id: u.Parent}
		if err := parent.Read(); err == nil {
			return parent
		}
	}
	return nil
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
	var user *User
	users := Users()

	if strings.IndexRune(usernameOrEmail, '@') == -1 {
		err := users.Filter("UserName", usernameOrEmail).One(user)
		return user, err
	}

	err := users.Filter("Email", usernameOrEmail).One(user)
	return user, err
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
			user.Role = UserRoleSuperAdmin
			user.IsForbid = false
			err = user.Update("Role", "IsForbid")
		}
	}
	return err
}
