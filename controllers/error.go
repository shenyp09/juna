package controllers

import (
    _ "github.com/astaxie/beego"
)

type ErrorController struct {
    baseRouter
}

func (c *ErrorController) Error404() {
    c.Data["content"] = "page not found"
    c.Layout = "layout.html"
    c.TplNames = "404.tpl"
    c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "html_head.tpl"
    c.LayoutSections["Header"] = "header.tpl"
    c.LayoutSections["Footer"] = "footer.tpl"
    c.LayoutSections["Scripts"] = "scripts.tpl"
}
