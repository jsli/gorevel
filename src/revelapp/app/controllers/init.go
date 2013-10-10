package controllers

import (
	"fmt"
	"github.com/robfig/revel"
)

func init() {
	revel.OnAppStart(func() {
		uploadPath = fmt.Sprintf("%s/public/upload/", revel.BasePath)
	})

	revel.InterceptMethod((*Application).adding, revel.BEFORE)
}
