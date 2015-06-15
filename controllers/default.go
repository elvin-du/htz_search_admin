package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
    c.Data["kinds"] = []string{"经典","学生问答","心得体会"}
	c.TplNames = "index.html"
}
