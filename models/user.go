package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      int
	Name    string
	Profile string
	Age     int
}

const (
	_MYSQL_DRIVER = "mysql"
)

func init() {
	orm.RegisterModel(new(User), new(DUser))
}

func Re() {

	orm.RegisterDriver(_MYSQL_DRIVER, orm.DRMySQL)

	orm.RegisterDataBase("default", _MYSQL_DRIVER, "root:root@/dp?charset=utf8")
}
