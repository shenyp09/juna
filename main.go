package main

import (
    "github.com/astaxie/beego"
    "github.com/beego/i18n"
    "html/template"
    "juna/controllers"
    _ "juna/routers"
)

type User struct {
    Id    int
    Name  string
    Count int
}

func main() {
    beego.SessionOn = true
    beego.ErrorController(&controllers.ErrorController{})
    beego.AddFuncMap("i18n", i18n.Tr)

    beego.EnableAdmin = true // Port 8088

    beego.SetLogger("file", `{"filename":"logs/juna.log"}`)
    beego.SetLevel(beego.LevelInformational)
    beego.SetLogFuncCall(true)

    beego.AddFuncMap("news", news)
    beego.Run()
}

func news(in string) (out template.HTML) {
    out = beego.Str2html(in)
    return
}
