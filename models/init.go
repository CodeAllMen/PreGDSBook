/**
  create by yy on 2019-08-21
*/

package models

import "github.com/astaxie/beego/orm"

func init() {
	//orm.Debug = true
	orm.RegisterModel(new(BookModel), new(CategoryModel))
}
