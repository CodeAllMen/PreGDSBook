/**
  create by yy on 2020/6/16
*/

package models

import "github.com/astaxie/beego/orm"

func RegistereUser(user Users) {
	o := orm.NewOrm()
	o.Insert(&user)
}

