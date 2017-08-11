package routers

import (
	"dpweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/catalog", &controllers.CatalogController{})
	beego.Router("/topic", &controllers.TopicController{})
}
