package controllers

import (
    _ "github.com/astaxie/beego"
    "strconv"
    "strings"
)

type NewsController struct {
    baseRouter
}

func (c *NewsController) Get() {

    c.Layout = "layout.html"
    c.TplNames = "news.tpl"
    c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "html_head.tpl"
    c.LayoutSections["Header"] = "header.tpl"
    c.LayoutSections["Footer"] = "footer.tpl"
    c.LayoutSections["Scripts"] = "scripts.tpl"

    v := XML_unmarshal(c.Lang)
    var newsAll string
    for i := 0; i < len(v.Nws); i++ {
        var tmp string
        newsID := len(v.Nws) - i
        tmp += "<li>"
        tmp += "<div class=\"newsList-all-item\">\n"
        date := strings.Split(v.Nws[i].Date, " ")
        tmp += "<span class=\"date\">" + date[0] + "/" + date[1] + "/" + date[2] + "</span>\n"
        tmp += "<a href=\"/object/news/" + strconv.Itoa(newsID) + "\" class=\"newsList-all-title\">" + v.Nws[i].Title + "</a>\n"
        tmp += "</div>\n</li>\n"
        newsAll += tmp
    }
    c.Data["Content"] = newsAll
}
