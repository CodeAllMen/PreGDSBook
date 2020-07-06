/**
  create by yy on 2020/3/27
*/

package sk

import "github.com/astaxie/beego"

type SubPage struct {
	beego.Controller
}

func (s *SubPage) Lp() {
	s.TplName = "lp/sk/lp.html"
}

func (s *SubPage) Terms() {
	s.TplName = "lp/sk/terms.html"
}

func (s *SubPage) Privacy() {
	s.TplName = "lp/sk/privacy.html"
}

func (s *SubPage) Help() {
	s.TplName = "lp/sk/help.html"
}
