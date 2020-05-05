/**
  create by yy on 2019-09-03
*/

package controllers

import "github.com/MobileCPX/PreGDSBook/models"

type ReadController struct {
	BaseController
}

func (r *ReadController) Read() {
	bookId, _ := r.GetInt("id")
	book, _ := models.GetBookById(bookId)
	r.Data["Book"] = book
	r.Data["Title"] = book.Title
	r.TplName = "read/read.html"
}
