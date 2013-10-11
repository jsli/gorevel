package models

import (
	"crypto/md5"
	"fmt"
	"github.com/coocood/qbs"
	"github.com/robfig/revel"
	"io"
	"regexp"
	"strings"
	"time"
)

type User struct {
	Id              int64
	Name            string      `qbs:"size:32"`
	Email           string      `qbs:"size:32"`
	Password        string      `qbs:"-"`
	ConfirmPassword string      `qbs:"-"`
	HashedPassword  string      `qbs:"size:32"`
	Type            int         `qbs:"default:2"` //1管理员，2普通用户
	Avatar          string      `qbs:"size:255"`
	Permissions     map[int]int `qbs:"-"`
	ValidateCode    string      `qbs:"size:255"`
	IsActive        bool
	Created         time.Time
	Updated         time.Time
}

var (
	nameRegex = regexp.MustCompile("^\\w*$")
)

func (user *User) Validate(v *revel.Validation) {
	v.Required(user.Name).Message("请输入用户名")
	valid := v.Match(user.Name, nameRegex).Message("只能使用字母、数字和下划线")
	if valid.Ok {
		if user.HasName() {
			err := &revel.ValidationError{
				Message: "该用户名已经注册过",
				Key:     "user.Name",
			}
			valid.Error = err
			valid.Ok = false

			v.Errors = append(v.Errors, err)
		}
	}

	v.Required(user.Email).Message("请输入Email")
	valid = v.Email(user.Email).Message("无效的电子邮件")
	if valid.Ok {
		if user.HasEmail() {
			err := &revel.ValidationError{
				Message: "该邮件已经注册过",
				Key:     "user.Email",
			}
			valid.Error = err
			valid.Ok = false

			v.Errors = append(v.Errors, err)
		}
	}

	v.Required(user.Password).Message("请输入密码")
	v.MinSize(user.Password, 3).Message("密码最少三位")
	v.Required(user.ConfirmPassword == user.Password).Message("密码不一致")
}

func (u *User) HasName() bool {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	user := new(User)
	q.WhereEqual("name", u.Name).Find(user)

	if user.Id > 0 {
		return true
	}
	return false

}

func (u *User) HasEmail() bool {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	user := new(User)
	q.WhereEqual("email", u.Email).Find(user)

	if user.Id > 0 {
		return true
	}
	return false
}

// 加密密码,转成md5
func EncryptPassword(password string) string {
	h := md5.New()
	io.WriteString(h, password)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (u *User) Save() bool {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer q.Close()

	if u.Password != "" {
		u.HashedPassword = EncryptPassword(u.Password)
	}

	_, err = q.Save(u)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func CheckSignin(name, password string) *User {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	user := new(User)
	condition := qbs.NewCondition("name = ?", name).And(
		"hashed_password = ?", EncryptPassword(password))
	q.Condition(condition).Find(user)

	return user
}

func FindUserById(id int64) *User {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	user := new(User)
	err = q.WhereEqual("id", id).Find(user)

	return user
}

func FindUserByCode(code string) *User {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	user := new(User)
	err = q.WhereEqual("ValidateCode", code).Find(user)

	return user
}

func (u *User) GetPermissions() map[int]int {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	if u.Permissions == nil {
		u.Permissions = make(map[int]int)
		var permissions []*Permissions

		err = q.WhereEqual("user_id", u.Id).FindAll(&permissions)

		for _, perm := range permissions {
			u.Permissions[perm.Perm] = perm.Perm
		}
	}

	return u.Permissions
}

func (u *User) IsAdmin() bool {
	return u.Type == 1
}

// 是否是默认头像
func (u *User) IsDefaultAvatar(avatar string) bool {
	return avatar == u.Avatar
}

// 头像的图片地址
func (u *User) AvatarImgSrc() string {
	if strings.HasPrefix(u.Avatar, "thumb") {
		return fmt.Sprintf("/public/upload/%s", u.Avatar)
	}
	return fmt.Sprintf("/public/img/%s", u.Avatar)
}
