/**
  create by yy on 2019-09-03
*/

package controllers

import (
	"github.com/astaxie/beego"

	"github.com/MobileCPX/PreGDSBook/models"
)

type HeaderStruct struct {
	Head []*models.CategoryModel
}

type BaseController struct {
	beego.Controller
}

func (b *BaseController) Prepare() {
	header := &HeaderStruct{}
	header.Head = models.GetCategoryList()
	b.Data["Header"] = header.Head
}
