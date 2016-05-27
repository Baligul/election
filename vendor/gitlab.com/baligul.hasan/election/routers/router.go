package routers

import (
	"gitlab.com/baligul.hasan/election/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/election/view", &controllers.ElectionController{}, "get:View")
	beego.Router("/election/home", &controllers.ElectionController{}, "*:Home")    
}
