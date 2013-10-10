package controllers

import (
	"fmt"
	"github.com/coocood/qbs"
	"github.com/robfig/revel"
	"revelapp/app/models"
	"revelapp/app/routes"
)

type Admin struct {
	Application
}

func (c *Admin) Index() revel.Result {
	return c.Render()
}

func (c *Admin) ListUser() revel.Result {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	var users []*models.User
	err = q.FindAll(&users)

	return c.Render(users)
}

func (c *Admin) DeleteUser(id int64) revel.Result {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	user := new(models.User)
	user.Id = id
	q.Find(user)

	_, err = q.Delete(user)
	if err != nil {
		fmt.Println(err)
	}

	return c.RenderJson([]byte("true"))
}

func (c *Admin) ListCategory() revel.Result {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	var categories []*models.Category
	err = q.FindAll(&categories)

	return c.Render(categories)
}

func (c *Admin) DeleteCategory(id int64) revel.Result {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	category := new(models.Category)
	category.Id = id
	q.Find(category)

	_, err = q.Delete(category)
	if err != nil {
		fmt.Println(err)
	}

	return c.RenderJson([]byte("true"))
}

func (c *Admin) NewCategory() revel.Result {
	title := "新建分类"
	return c.Render(title)
}

func (c *Admin) NewCategoryPost(category models.Category) revel.Result {
	category.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Admin.NewCategory())
	}

	if !category.Save() {
		c.Flash.Error("添加分类失败")
	}

	return c.Redirect(routes.Admin.ListCategory())
}

func (c *Admin) EditCategory(id int64) revel.Result {
	title := "编辑分类"

	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	category := new(models.Category)
	category.Id = id
	q.Find(category)

	c.Render(title, category)

	return c.RenderTemplate("Admin/NewCategory.html")
}

func (c *Admin) EditCategoryPost(id int64, category models.Category) revel.Result {
	category.Id = id
	category.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Admin.NewCategory())
	}

	if !category.Save() {
		c.Flash.Error("编辑分类失败")
	}

	return c.Redirect(routes.Admin.ListCategory())
}
