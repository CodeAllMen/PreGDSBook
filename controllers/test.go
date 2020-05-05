/**
  create by yy on 2019-09-05
*/

package controllers

type TestController struct {
	BaseController
}

func (t *TestController) Index() {
	t.TplName = "test/test.html"
}
