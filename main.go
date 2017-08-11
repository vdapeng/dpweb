package main

import (
	"dpweb/models"
	_ "dpweb/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.Re()
}

func main() {
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
