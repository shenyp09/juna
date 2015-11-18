package controllers

import (
    // "fmt"
    _ "github.com/astaxie/beego"
)

type AboutController struct {
    baseRouter
}

func (c *AboutController) Get() {

    path := c.Ctx.Input.Param(":path")
    if path == "" || path == "general" {
        c.Layout = "layout_sidebar.html"
        c.TplNames = "general.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["Sidebar"] = "sidebar_about.tpl"
        c.LayoutSections["HtmlHead"] = "html_head.tpl"
        c.LayoutSections["Header"] = "header.tpl"
        c.LayoutSections["Footer"] = "footer.tpl"
        c.LayoutSections["Scripts"] = "scripts.tpl"
        c.Data["GeneralActive"] = "active"
    } else if path == "collaboration" {
        c.Layout = "layout_sidebar.html"
        c.TplNames = "collaboration.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["Sidebar"] = "sidebar_about.tpl"
        c.LayoutSections["HtmlHead"] = "html_head.tpl"
        c.LayoutSections["Header"] = "header.tpl"
        c.LayoutSections["Footer"] = "footer.tpl"
        c.LayoutSections["Scripts"] = "scripts.tpl"
        c.Data["CollaborationActive"] = "active"
    } else {
        c.Abort("404")
    }

}
