/**
  create by yy on 2019-09-06
*/

package models

import "github.com/astaxie/beego/orm"

type CategoryModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (c *CategoryModel) TableName() string {
	return "category_en"
}

func GetCategoryList() []*CategoryModel {
	var categoryList []*CategoryModel
	o := orm.NewOrm()
	_, _ = o.QueryTable("category_en").All(&categoryList)
	return categoryList
}

func GetCategoryById(categoryId int) *CategoryModel {
	var category CategoryModel
	o := orm.NewOrm()
	_ = o.QueryTable("category_en").Filter("id", categoryId).One(&category)
	return &category
}
