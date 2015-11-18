package controllers

import (
    _ "github.com/astaxie/beego"
    "strconv"
    "strings"
)

type MainController struct {
    baseRouter
}

func (c *MainController) Get() {

    c.Layout = "layout.html"
    c.TplNames = "index.tpl"
    c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "html_head.tpl"
    c.LayoutSections["Header"] = "header.tpl"
    c.LayoutSections["Footer"] = "footer.tpl"
    c.LayoutSections["Scripts"] = "scripts.tpl"

    v := XML_unmarshal(c.Lang)
    c.Data["Website"] = "esic.ciae.ac.cn/juna"
    c.Data["Email"] = "ypshen@esic.ac.cn"

    c.Data["banner1_img"] = "/static/img/banner/1.jpg"
    c.Data["banner1_caption"] = "JUNA重大项目启动会召开"
    c.Data["banner1_link"] = "/object/news/4"
    c.Data["banner2_img"] = "/static/img/banner/2.jpg"
    c.Data["banner2_caption"] = "美国《科学》杂志报道原子能院核天体物理实验计划"
    c.Data["banner2_link"] = "/object/news/2"
    c.Data["banner3_img"] = "/static/img/banner/3.jpg"
    c.Data["banner3_caption"] = "核物理领域首个重大基金项目启动"
    c.Data["banner3_link"] = "/object/news/1"

    for i := 0; i < 3; i++ {
        date := strings.Split(v.Nws[i].Date, " ")
        stri := strconv.Itoa(i + 1)
        c.Data["news"+stri+"_day"] = date[0]
        c.Data["news"+stri+"_month"] = date[1]
        c.Data["news"+stri+"_year"] = date[2]
        c.Data["news"+stri+"_link"] = "/object/news/" + v.Nws[i].ID
        c.Data["news"+stri+"_title"] = v.Nws[i].Title
        c.Data["news"+stri+"_summary"] = v.Nws[i].Description
    }
    j := 0
    for i := 0; i < 6; i++ {
        for v.Nws[j].Focus != "1" {
            j += 1
        }
        stri := strconv.Itoa(i + 1)
        c.Data["focus"+stri+"_link"] = "/object/news/" + v.Nws[i].ID
        if v.Nws[i].Image != "" {
            c.Data["focus"+stri+"_img"] = "static/img/" + v.Nws[i].Image
        }
        c.Data["focus"+stri+"_title"] = v.Nws[i].Title
        c.Data["focus"+stri+"_summary"] = v.Nws[i].Description
        j++
    }
}
