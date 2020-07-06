/**
  create by yy on 2020/6/16
*/

package models

import "github.com/astaxie/beego/orm"

// CheckUser 检查用户是否存在
func CheckUser(userName string) (userID string) {
	o := orm.NewOrm()
	var user Users

	o.QueryTable("users").Filter("user_name", userName).One(&user)
	if user.Id != 0 {
		userID = user.UserName
	}
	return
}
