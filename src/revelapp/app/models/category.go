package models

import (
	"fmt"
	"github.com/coocood/qbs"
	"github.com/robfig/revel"
	"time"
)

type Category struct {
	Id      int64
	Name    string `qbs:"size:32"`
	Intro   string `qbs:"size:255"`
	Created time.Time
}

func (category *Category) Validate(v *revel.Validation) {
	valid := v.Required(category.Name).Message("请输入名称")

	if valid.Ok {
		if category.HasName() {
			err := &revel.ValidationError{
				Message: "该名称已存在",
				Key:     "category.Name",
			}
			valid.Error = err
			valid.Ok = false

			v.Errors = append(v.Errors, err)
		}
	}
}

func (c *Category) HasName() bool {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	category := new(Category)
	condition := qbs.NewCondition("name = ?", c.Name)
	if c.Id > 0 {
		condition = qbs.NewCondition("name = ?", c.Name).And("id != ?", c.Id)
	}
	err = q.Condition(condition).Find(category)

	if category.Id > 0 {
		return true
	}
	return false

}

func (c *Category) Save() bool {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer q.Close()

	_, err = q.Save(c)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetCategories() []*Category {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	var categories []*Category
	err = q.FindAll(&categories)
	if err != nil {
		fmt.Println(err)
	}

	return categories
}
