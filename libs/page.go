/**
  create by yy on 2019-09-03
*/

package libs

type Page struct {
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Data  interface{} `json:"data"`
}

/**
包装分页数据返回
*/
func GetPageData(data interface{}, page *int, total int, size int) *Page {
	return &Page{
		Total: total / size,
		Page:  *page,
		Data:  data,
	}
}

/**
获取分页时的起始位置
*/
func GetOffset(page int, size int) int {
	return (page - 1) * size
}
