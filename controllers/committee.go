package controllers

import (
    _ "github.com/astaxie/beego"
)

type CommitteeController struct {
    baseRouter
}

func (c *CommitteeController) Get() {

    path := c.Ctx.Input.Param(":path")
    if path == "" || path == "iac" {
        c.Layout = "layout_sidebar.html"
        c.TplNames = "IAC.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = "html_head.tpl"
        c.LayoutSections["Header"] = "header.tpl"
        c.LayoutSections["Footer"] = "footer.tpl"
        c.LayoutSections["Scripts"] = "scripts.tpl"
        c.LayoutSections["Sidebar"] = "sidebar_committee.tpl"
        c.Data["IACActive"] = "active"
    } else {
        c.Abort("404")
    }
}
