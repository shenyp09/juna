package controllers

import (
    _ "github.com/astaxie/beego"
)

type PubController struct {
    baseRouter
}

func (c *PubController) Get() {

    path := c.Ctx.Input.Param(":path")
    if path == "" || path == "publications" {
        c.Layout = "layout_sidebar.html"
        c.TplNames = "pub/publications.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = "html_head.tpl"
        c.LayoutSections["Header"] = "header.tpl"
        c.LayoutSections["Footer"] = "footer.tpl"
        c.LayoutSections["Scripts"] = "scripts.tpl"
        c.LayoutSections["Sidebar"] = "sidebar_publications.tpl"
        c.Data["PublicationsActive"] = "active"
    } else if path == "proposals" {
        c.Layout = "layout_sidebar.html"
        c.TplNames = "pub/proposals.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = "html_head.tpl"
        c.LayoutSections["Header"] = "header.tpl"
        c.LayoutSections["Footer"] = "footer.tpl"
        c.LayoutSections["Scripts"] = "scripts.tpl"
        c.LayoutSections["Sidebar"] = "sidebar_publications.tpl"
        c.Data["ProposalsActive"] = "active"
    } else if path == "talks" {
        c.Layout = "layout_sidebar.html"
        c.TplNames = "pub/talks.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = "html_head.tpl"
        c.LayoutSections["Header"] = "header.tpl"
        c.LayoutSections["Footer"] = "footer.tpl"
        c.LayoutSections["Scripts"] = "scripts.tpl"
        c.LayoutSections["Sidebar"] = "sidebar_publications.tpl"
        c.Data["TalksActive"] = "active"
    } else if path == "theses" {
        c.Layout = "layout_sidebar.html"
        c.TplNames = "pub/theses.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = "html_head.tpl"
        c.LayoutSections["Header"] = "header.tpl"
        c.LayoutSections["Footer"] = "footer.tpl"
        c.LayoutSections["Scripts"] = "scripts.tpl"
        c.LayoutSections["Sidebar"] = "sidebar_publications.tpl"
        c.Data["ThesesActive"] = "active"
    } else {
        c.Abort("404")
    }
}
