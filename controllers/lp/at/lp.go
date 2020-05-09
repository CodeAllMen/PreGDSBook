/**
  create by yy on 2020/5/5
*/

package at

import "github.com/astaxie/beego"

type SubPage struct {
	beego.Controller
}

func (s *SubPage) Lp() {
	s.TplName = "lp/at/lp.html"
}

func (s *SubPage) Tan() {
	s.TplName = "at/tan.html"
}

func (s *SubPage) Confirm() {
	s.TplName = "at/confirm.html"
}

func (s *SubPage) Condition() {
	s.TplName = "at/tnc.html"
}

func (s *SubPage) Help() {
	s.TplName = "at/help.html"
}

func (s *SubPage) Privacy() {
	s.TplName = "at/privacy.html"
}

func (s *SubPage) AGB() {
	s.TplName = "lp/at/AGB.html"
}

func (s *SubPage) Impressum() {
	s.TplName = "lp/at/IMPRESSUM.html"
}

func (s *SubPage) Datenschutz() {
	s.TplName = "lp/at/Datenschutz.html"
}

func (s *SubPage) KONTAKT() {
	s.TplName = "lp/at/KONTAKT.html"
}

func (s *SubPage) KUNDIGUNG() {
	s.TplName = "lp/at/KUNDIGUNG.html"
}

func (s *SubPage) Rucktrittsrechts() {
	s.TplName = "lp/at/Rucktrittsrechts.html"
}

func (s *SubPage) FAQ() {
	s.TplName = "lp/at/FAQ.html"
}

func (s *SubPage) Welcome() {
	s.TplName = "lp/at/welcome.html"
}
