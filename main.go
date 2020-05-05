package main

import (
	_ "github.com/MobileCPX/PreGDSBook/init"
	_ "github.com/MobileCPX/PreGDSBook/routers"
	_ "github.com/MobileCPX/PreGDSBook/template_func"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
