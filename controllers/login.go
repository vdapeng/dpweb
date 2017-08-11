package controllers

import (
	"github.com/astaxie/beego"

	"dpweb/models"

	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	fmt.Println("Login")
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	error := models.AddDUser(username, password)
	if error != nil {
		beego.Error(error)
	}

	c.Redirect("/", 302)
	return
}
