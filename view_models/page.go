/**
  create by yy on 2019-09-11
*/

package view_models

import (
	"fmt"

	"github.com/MobileCPX/PreGDSBook/data_struct"
)

func PackagePage(page, size, total int, data interface{}) *data_struct.Page {
	fmt.Println(size)
	return &data_struct.Page{
		Page:      page,
		TotalPage: total / size,
		Data:      data,
	}
}
