/**
  create by yy on 2020/3/20
*/

package lp

import "github.com/astaxie/beego"

type SubPage struct {
	beego.Controller
}

func (s *SubPage) Lp() {
	s.TplName = "lp/ae/lp.html"
}

func (s *SubPage) Pin() {
	s.TplName = "lp/ae/pin.html"
}