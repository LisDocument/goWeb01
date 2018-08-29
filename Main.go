package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static/images","images")
	beego.SetStaticPath("/static/css","css")
	beego.SetStaticPath("/static/js","js")

	beego.Router("/login",&controllers.LoginController{})
	fmt.Print("11111")
}