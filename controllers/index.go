/**
  create by yy on 2019-09-03
*/

package controllers

import (
	"github.com/MobileCPX/PreGDSBook/models"
)

type IndexController struct {
	BaseController
}

func (index *IndexController) Get() {
	params := make(map[string]interface{})
	params["size"] = 12
	params["page"] = 2
	bookList, _ := models.GetBookList(params)
	params["page"] = 3
	completeBookList, _ := models.GetBookList(params)
	params["page"] = 1
	hotBookList, _ := models.GetBookList(params)
	index.Data["HotBookList"] = hotBookList
	index.Data["BookList"] = bookList
	index.Data["CompleteBookList"] = completeBookList
	index.Data["Title"] = "gdsbook"
	index.Layout = "common/layout.html"
	index.TplName = "index/index.html"
}

func (index *IndexController) Index() {
	params := make(map[string]interface{})
	params["size"] = 12
	params["page"] = 2
	bookList, _ := models.GetBookList(params)
	params["page"] = 3
	params["page"] = 1
	hotBookList, _ := models.GetBookList(params)
	index.Data["HotBookList"] = hotBookList
	completeBookList, _ := models.GetBookList(params)
	index.Data["BookList"] = bookList
	index.Data["Title"] = "gdsbook"
	index.Data["CompleteBookList"] = completeBookList
	index.Layout = "common/layout.html"
	index.TplName = "index/index.html"
}
