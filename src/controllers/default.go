package controllers

import (
	"strconv"

	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	// Obtain the values of the route parameters defined in the route above
	operation := c.Ctx.Input.Param(":operation")
	num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
	num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))

	// Set the values for use in the template
	c.Data["operation"] = operation
	c.Data["num1"] = num1
	c.Data["num2"] = num2
	c.TplName = "result.html"

	// Perform the calculation depending on the 'operation' route parameter
	switch operation {
	case "sum":
		c.Data["result"] = Add(num1, num2)
	case "product":
		c.Data["result"] = Multiply(num1, num2)
	default:
		c.TplName = "invalid-route.html"
	}
}

func Add(n1, n2 int) int {
	return n1 + n2
}

func Multiply(n1, n2 int) int {
	return n1 * n2
}
