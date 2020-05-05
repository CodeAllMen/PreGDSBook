/**
  create by yy on 2019-08-21
*/

package template_func

import "github.com/astaxie/beego"

func init()  {
	_ = beego.AddFuncMap("mod", mod)
}
