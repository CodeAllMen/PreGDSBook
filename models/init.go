/**
  create by yy on 2019-08-21
*/

package models

import "github.com/astaxie/beego/orm"

type Users struct {
	Id       int64 `orm:"pk;auto"`
	UserName string
	Password string
	Sp       string
	Country  string
}

func init() {
	// orm.Debug = true
	orm.RegisterModel(new(BookModel), new(CategoryModel), new(Users))
	orm.RunSyncdb("default", false, true)
}
