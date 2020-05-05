/**
  create by yy on 2019-09-09
*/

package controllers

import (
	"github.com/MobileCPX/PreGDSBook/models"
	"github.com/MobileCPX/PreGDSBook/view_models"
)

type SearchController struct {
	BaseController
}

func (s *SearchController) Search() {
	page, _ := s.GetInt("page")
	hot, _ := s.GetInt("hot")
	complete, _ := s.GetInt("completed")
	if page == 0 {
		page = 1
	}
	if hot == 1 && page == 1 {
		page = 5
	}
	if complete == 1 && page == 1 {
		page = 8
	}
	search := s.GetString("search")
	params := make(map[string]interface{})
	params["page"] = page
	size := 12
	params["size"] = 12
	bookList, total := models.GetBookListByTitle(params, search)
	data := view_models.PackagePage(page, size, total, bookList)
	s.Data["BookList"] = data
	s.Data["Title"] = search
	s.Data["Search"] = search
	s.Layout = "common/layout.html"
	s.TplName = "search/search.html"
}

func (s *SearchController) Category() {
	page, _ := s.GetInt("page")
	if page == 0 {
		page = 1
	}
	search, _ := s.GetInt("category")
	if search == 0 {
		search = 1
	}
	params := make(map[string]interface{})
	params["page"] = page
	size := 12
	params["size"] = 12
	bookList, total := models.GetBookListByCategory(params, search)
	category := models.GetCategoryById(search)
	data := view_models.PackagePage(page, size, total, bookList)
	s.Data["BookList"] = data
	s.Data["Category"] = category.Name
	s.Data["CategoryId"] = category.Id
	s.Data["Title"] = category.Name
	s.Layout = "common/layout.html"
	s.TplName = "search/category.html"
}
