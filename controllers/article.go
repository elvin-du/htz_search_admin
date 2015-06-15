package controllers

import (
	"htz_search_admin/models"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Post() {
	title := c.GetString("title")
	content := c.GetString("content")
	kind := c.GetString("kind")
	author := c.GetString("author")
	ctime := c.GetString("ctime")
	err := models.ArticleModel().Add(title, content, kind, author, ctime)
	if nil != err {
		c.Abort("500")
	}
	c.Data["kinds"] = []string{"经典", "学生问答", "心得体会"}
	c.TplNames = "index.html"
}
