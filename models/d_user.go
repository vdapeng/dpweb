package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	_D_USER = "DUser"
)

type DUser struct {
	Id       int
	Username string
	Password string
	Email    string
	Name     string
	Tel      string
	Created  time.Time
}

func AddDUser(username, password string) error {
	o := orm.NewOrm()
	du := &DUser{
		Username: username,
		Password: password,
		Created:  time.Now(),
	}
	id, err := o.Insert(du)
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}

func GetAll() ([]*DUser, error) {
	o := orm.NewOrm()
	users := make([]*DUser, 0)
	qs := o.QueryTable("d_user")
	_, err := qs.All(&users)
	return users, err
}
