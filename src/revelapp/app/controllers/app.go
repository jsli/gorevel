package controllers

import (
	"fmt"
	"github.com/coocood/qbs"
	"github.com/robfig/revel"
	"revelapp/app/models"
	"revelapp/app/routes"
)

type Application struct {
	*revel.Controller
}

func (c *Application) adding() revel.Result {
	c.RenderArgs["active"] = c.Name
	user := c.connected()
	if user != nil {
		c.RenderArgs["user"] = user
	}

	// 检查是否需要授权
	value, ok := Permissions[c.Action]
	if ok {
		if user == nil {
			c.Flash.Error("请先登录")
			c.Session["preUrl"] = c.Request.Request.URL.String()
			return c.Redirect(routes.User.Signin())
		} else {
			perm := user.GetPermissions()
			_, ok := perm[value]
			if !ok {
				return c.Forbidden("抱歉，您没有得到授权！")
			}
		}
	}
	return nil
}

func (c *Application) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username)
	}
	return nil
}

func (c *Application) getUser(username string) *models.User {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	user := new(models.User)
	q.WhereEqual("name", username).Find(user)

	if user.Id == 0 {
		return nil
	}

	return user
}

type App struct {
	Application
}

func (c *App) Index() revel.Result {
	t1 := `{{set . "title" "Home"}}
{{template "header.html" .}}

<header class="hero-unit" style="background-color:#A9F16C">
  <div class="container">
	<div class="row">
	  <div class="hero-text">
		<h1>It works!</h1>
		<p></p>
	  </div>
	</div>
  </div>
</header>

<div class="container">
  <div class="row">
	<div class="span6">
	  {{template "flash.html" .}}
	</div>
  </div>
</div>

{{template "footer.html" .}}`

	return c.Render(t1)
}
