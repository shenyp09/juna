package routers

import (
    "github.com/astaxie/beego"
    "github.com/beego/i18n"
    "juna/controllers"

    "strings"
)

// langType represents a language type.
type langType struct {
    Lang, Name string
}

func init() {
    beego.SessionCookieLifeTime = 60

    // Initialized language type list.
    langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
    names := strings.Split(beego.AppConfig.String("lang::names"), "|")
    langTypes := make([]*langType, 0, len(langs))
    for i, v := range langs {
        langTypes = append(langTypes, &langType{
            Lang: v,
            Name: names[i],
        })
    }

    for _, lang := range langs {
        beego.Trace("Loading language: " + lang)
        if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
            beego.Error("Fail to set message file: " + err.Error())
            return
        }
    }

    beego.Router("/", &controllers.MainController{})
    beego.Router("/index", &controllers.MainController{})
    beego.Router("/about/?:path", &controllers.AboutController{})
    beego.Router("/news", &controllers.NewsController{})
    beego.Router("/committee/?:path", &controllers.CommitteeController{})
    beego.Router("/pub/?:path", &controllers.PubController{})
    beego.AutoRouter(&controllers.ObjectController{})
}
