/**
  create by yy on 2019-09-03
*/

package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/guaidashu/go_helper"
)

type BookModel struct {
	Id              int    `json:"id"`
	Author          string `json:"author"`
	BookUrl         string `json:"book_url"`
	Title           string `json:"title"`
	Language        string `json:"language"`
	UniformTitle    string `json:"uniform_title"`
	Translator      string `json:"translator"`
	LocClass        string `json:"loc_class"`
	Subject         string `json:"subject"`
	Category        string `json:"category"`
	EbookNo         string `json:"ebook_no"`
	ReleaseDate     string `json:"release_date"`
	CopyrightStatus string `json:"copyright_status"`
	DownloadStatus  string `json:"download_status"`
	Price           string `json:"price"`
	Note            string `json:"note"`
	Editor          string `json:"editor"`
	Contributor     string `json:"contributor"`
	SourceUrl       string `json:"source_url"`
	Illustrator     string `json:"illustrator"`
	Commentator     string `json:"commentator"`
	Annotator       string `json:"annotator"`
	Contents        string `json:"contents"`
	LocNo           string `json:"loc_no"`
	AuthorIntroduce string `json:"author_introduce"`
	Imprint         string `json:"imprint"`
	Compiler        string `json:"compiler"`
	AlternateTitle  string `json:"alternate_title"`
	SeriesTitle     string `json:"series_title"`
	CategoryId      string `json:"category_id"`
}

func (index *BookModel) TableName() string {
	return "book_en"
}

func GetBookById(id int) (*BookModel, error) {
	var bookModel BookModel
	o := orm.NewOrm()
	err := o.QueryTable("book_en").Filter("id", id).One(&bookModel)
	if err != nil {
		return nil, err
	}
	return &bookModel, nil
}

/**
获取电子书列表
*/
func GetBookList(params map[string]interface{}) ([]*BookModel, error) {
	var bookList []*BookModel
	var page int
	var limit int
	if item, ok := params["page"]; ok {
		page = item.(int)
	}
	if item, ok := params["size"]; ok {
		limit = item.(int)
	} else {
		limit = 12
	}
	limit = params["size"].(int)
	offset := go_helper.GetStartPos(page, limit)
	o := orm.NewOrm()
	db := o.QueryTable("book_en")
	_, err := db.Limit(limit, offset).All(&bookList)
	return bookList, err
}

/**
获取电子书列表 根据标题
*/
func GetBookListByTitle(params map[string]interface{}, title string) ([]*BookModel, int) {
	var bookList []*BookModel
	var page int
	var limit int
	if item, ok := params["page"]; ok {
		page = item.(int)
	}
	if item, ok := params["size"]; ok {
		limit = item.(int)
	} else {
		limit = 12
	}
	limit = params["size"].(int)
	offset := go_helper.GetStartPos(page, limit)
	o := orm.NewOrm()
	db := o.QueryTable("book_en").Filter("title__icontains", title)
	num, _ := db.Count()
	_, _ = db.Limit(limit, offset).All(&bookList)
	return bookList, int(num)
}

/**
获取电子书列表 根据标题
*/
func GetBookListByCategory(params map[string]interface{}, categoryId int) ([]*BookModel, int) {
	var bookList []*BookModel
	var page int
	var limit int
	if item, ok := params["page"]; ok {
		page = item.(int)
	}
	if item, ok := params["size"]; ok {
		limit = item.(int)
	} else {
		limit = 12
	}
	limit = params["size"].(int)
	offset := go_helper.GetStartPos(page, limit)
	o := orm.NewOrm()
	db := o.QueryTable("book_en").Filter("category_id", categoryId)
	num, _ := db.Count()
	_, _ = db.Limit(limit, offset).All(&bookList)
	return bookList, int(num)
}
