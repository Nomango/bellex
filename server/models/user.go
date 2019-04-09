package models

import (
	"errors"
	"fmt"
	"regexp"
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
	NickName string `orm:"size(30)" json:"nickname"`
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

// VerifyPassword compare raw password and encoded password
func VerifyPassword(rawPwd, encodedPwd string) bool {

	// split
	var salt, encoded string
	if len(encodedPwd) > 11 {
		salt = encodedPwd[:10]
		encoded = encodedPwd[11:]
	}

	return utils.EncodePassword(rawPwd, salt) == encoded
}

// HasUser returns whether if user exists
func HasUser(usernameOrEmail string) bool {
	users := Users()
	if strings.IndexRune(usernameOrEmail, '@') == -1 {
		return users.Filter("UserName", usernameOrEmail).Exist()
	}
	return users.Filter("Email", strings.ToLower(usernameOrEmail)).Exist()
}

// FindUser ...
func FindUser(usernameOrEmail string) (*User, error) {
	var user User
	users := Users()

	if strings.IndexRune(usernameOrEmail, '@') == -1 {
		err := users.Filter("UserName", usernameOrEmail).One(&user)
		return &user, err
	}

	err := users.Filter("Email", strings.ToLower(usernameOrEmail)).One(&user)
	return &user, err
}

func CheckRegister(user *User) error {
	if Users().Filter("UserName", user.UserName).Exist() {
		return errors.New("账号已存在")
	}

	user.Email = strings.ToLower(user.Email)
	emailReg := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	if matched, _ := regexp.MatchString(emailReg, user.Email); !matched {
		return errors.New("邮箱格式不合法")
	}

	if len(user.UserName) < 3 || len(user.Password) < 4 {
		return errors.New("用户名或密码不合法")
	}
	return nil
}

// RegisterUser ...
func RegisterUser(user *User) error {
	if err := CheckRegister(user); err != nil {
		return nil
	}

	user.SetNewPassword(user.Password)
	return user.Insert()
}

// CreateAdminUser register an administrator
func CreateAdminUser() error {
	if !HasUser("admin") {
		user := &User{
			UserName:    "admin",
			Email:       "admin@localhost.com",
			Password:    "admin",
			Role:        UserRoleAdmin,
			IsForbid:    false,
			Institution: &Institution{Id: 1},
		}
		return RegisterUser(user)
	}
	return nil
}
