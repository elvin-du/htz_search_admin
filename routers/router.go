package routers

import (
	"github.com/astaxie/beego"
	"htz_search_admin/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/article", &controllers.ArticleController{})
}
