package main

import (
	"mathapp/controllers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	/* This would match routes like the following
	   /sum/3/5
	   /product/6/23
	   ...
	*/
	web.Router("/:operation/:num1:int/:num2:int", &controllers.MainController{})
	web.Run()
}
