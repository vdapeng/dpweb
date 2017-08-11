package controllers

import (
	"github.com/astaxie/beego"
)

type CatalogController struct {
	beego.Controller
}

func (c *CatalogController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["IsCatalog"] = true
	c.TplName = "catalog.html"
}
