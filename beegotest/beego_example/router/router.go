package router

import (
	"github.com/astaxie/beego"
	"github.com/ggvylf/learngo/projects/beego_example/controller/IndexController"
)

func init() {
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")
}
