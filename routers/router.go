package routers

import (
	"github.com/Baligul/election/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/election/view", &controllers.ElectionController{}, "get:View")
	beego.Router("/election/home", &controllers.ElectionController{}, "*:Home")    
}
