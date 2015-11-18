package controllers

import (
    "fmt"
    _ "github.com/astaxie/beego"
    "strconv"
    "strings"
)

type ObjectController struct {
    baseRouter
}

func (c *ObjectController) News() {
    session_controller(&c.Controller)

    c.Layout = "layout.html"
    c.TplNames = "news_item.tpl"
    c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "html_head.tpl"
    c.LayoutSections["Header"] = "header.tpl"
    c.LayoutSections["Footer"] = "footer.tpl"
    c.LayoutSections["Scripts"] = "scripts.tpl"

    newsID, err := strconv.Atoi(c.Ctx.Input.Params["0"])
    if err != nil {
        fmt.Println(err)
        c.Abort("404")
    }
    v := XML_unmarshal(c.Lang)
    newsID = len(v.Nws) - newsID
    if newsID < len(v.Nws) {
        date := strings.Split(v.Nws[newsID].Date, " ")
        c.Data["news_day"] = date[0]
        c.Data["news_month"] = date[1]
        c.Data["news_year"] = date[2]
        c.Data["news_link"] = ""
        if v.Nws[newsID].Image != "" {
            c.Data["news_img"] = "/static/img/" + v.Nws[newsID].Image
        }
        c.Data["news_title"] = v.Nws[newsID].Title
        c.Data["news_summary"] = v.Nws[newsID].Description
        c.Data["news_content"] = v.Nws[newsID].Content.Description
    } else {
        c.Abort("404")
    }

}
