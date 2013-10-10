package controllers

import (
	"github.com/disintegration/imaging"
	"github.com/robfig/revel"
	"image"
	"revelapp/app/models"
	"revelapp/app/routes"
	"strings"
)

var (
	avatars = []string{
		"gopher_teal.jpg",
		"gopher_aqua.jpg",
		"gopher_brown.jpg",
		"gopher_strawberry_bg.jpg",
		"gopher_strawberry.jpg",
	}
	defaultAvatar = avatars[0]
)

type User struct {
	Application
}

func (c *User) Signup() revel.Result {
	return c.Render()
}

func (c *User) SignupPost(user models.User) revel.Result {
	user.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.User.Signup())
	}
	user.Type = MemberGroup
	user.Avatar = defaultAvatar
	if !user.Save() {
		c.Flash.Error("注册用户失败")
		return c.Redirect(routes.User.Signup())
	}

	perm := new(models.Permissions)
	perm.UserId = user.Id
	perm.Perm = MemberGroup
	perm.Save()

	c.Session["user"] = user.Name

	return c.Redirect(routes.App.Index())
}

func (c *User) Signin() revel.Result {
	return c.Render()
}

func (c *User) SigninPost(name, password string) revel.Result {
	c.Validation.Required(name).Message("请输入用户名")
	c.Validation.Required(password).Message("请输入密码")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.User.Signin())
	}

	user := models.CheckSignin(name, password)

	if user.Id == 0 {
		c.Validation.Keep()
		c.FlashParams()
		c.Flash.Out["user"] = name
		c.Flash.Error("用户名或密码错误")
		return c.Redirect(routes.User.Signin())
	}

	c.Session["user"] = name

	preUrl, ok := c.Session["preUrl"]
	if ok {
		return c.Redirect(preUrl)
	}

	return c.Redirect(routes.App.Index())
}

func (c *User) Signout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}

	return c.Redirect(routes.App.Index())
}

func (c *User) Edit(id int64) revel.Result {
	user := models.FindUserById(id)
	if user.Id == 0 {
		return c.NotFound("用户不存在")
	}

	return c.Render(user, avatars)
}

func (c *User) EditPost(id int64, avatar string) revel.Result {
	checkFileExt(c.Controller, imageExts, "picture", "Only image")
	user := models.FindUserById(id)
	if user.Id == 0 {
		return c.NotFound("用户不存在")
	}

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.User.Edit(id))
	}

	if ok, _ := getFileExt(c.Request, "picture"); ok {
		picture := saveFile(c.Request, "picture")
		src, _ := imaging.Open(uploadPath + picture)
		var dst *image.NRGBA

		dst = imaging.Thumbnail(src, 48, 48, imaging.CatmullRom)
		avatar = "thumb" + picture
		imaging.Save(dst, uploadPath+avatar)
		deleteFile(picture)
	}

	if avatar != "" {
		if strings.HasPrefix(user.Avatar, "thumb") {
			deleteFile(user.Avatar)
		}
		user.Avatar = avatar
	}

	if user.Save() {
		c.Flash.Success("保存成功")
	} else {
		c.Flash.Error("保存信息失败")
	}

	return c.Redirect(routes.User.Edit(id))
}
