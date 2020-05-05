/**
  create by yy on 2020/3/23
*/

package mls

import "github.com/astaxie/beego"

type SubPage struct {
	beego.Controller
}

func (s *SubPage) Lp() {
	s.TplName = "lp/mls/lp.html"
}

func (s *SubPage) Pin() {
	s.TplName = "lp/mls/pin.html"
}

func (s *SubPage) Thank() {
	s.TplName = "lp/mls/thank.html"
}