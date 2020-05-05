/**
  create by yy on 2019-09-11
*/

package data_struct

type Page struct {
	Page      int         `json:"page"`
	TotalPage int         `json:"total"`
	Data      interface{} `json:"data"`
}
