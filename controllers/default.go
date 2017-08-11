package controllers

import (
	"dpweb/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "home.html"
	var err error
	c.Data["Users"], err = models.GetAll()
	if err != nil {
		beego.Error(err)
	}
}
