package routers

import (
	"github.com/MobileCPX/PreGDSBook/controllers/lp"
	"github.com/MobileCPX/PreGDSBook/controllers/lp/mls"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/MobileCPX/PreGDSBook/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.ReadController{})
	beego.AutoRouter(&controllers.TestController{})
	beego.AutoRouter(&controllers.SearchController{})

	beego.Router("/lp", &lp.SubPage{}, "*:Lp")
	beego.Router("/pin", &lp.SubPage{}, "*:Pin")

	beego.Router("/mls/lp", &mls.SubPage{}, "*:Lp")
	beego.Router("/mls/pin", &mls.SubPage{}, "*:Pin")
	beego.Router("/mls/thank", &mls.SubPage{}, "*:Thank")

	// beego.Router("/lp", &controllers.SubFlowController{})
	// beego.Router("/help", &controllers.SubFlowController{}, "Get:ServiceManagement")
	// beego.Router("/compat", &controllers.SubFlowController{}, "Get:Compat")
	// beego.Router("/tnc", &controllers.SubFlowController{}, "Get:TNC")
	// beego.Router("/info", &controllers.SubFlowController{}, "Get:InfoPrivacy")

	beego.ErrorHandler("404", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Not Found"))
		if err != nil {
			logs.Error(err)
		}
	})

	beego.ErrorHandler("500", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("internal error"))
		if err != nil {
			logs.Error(err)
		}
	})
}
