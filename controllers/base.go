package controllers

import (
    "github.com/astaxie/beego"
    "github.com/beego/i18n"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
    "strings"
    "time"
)

type UV struct {
    Name  string
    Count int
}

type Userinfo struct {
    Mark          int64
    PageStartTime string
    IP            string
    Lang          string
    UserAgent     string
}

type NestPreparer interface {
    NestPrepare()
}

var langTypes []*langType // Languages are supported.

// langType represents a language type.
type langType struct {
    Lang, Name string
}

// baseRouter implemented global settings for all other routers.
type baseRouter struct {
    beego.Controller
    i18n.Locale
    isNew bool
}

// Prepare implemented Prepare method for baseRouter.
func (c *baseRouter) Prepare() {
    // Setting properties.

    c.Data["PageStartTime"] = time.Now()

    // Redirect to make URL clean.
    if c.setLangVer() {
        i := strings.Index(c.Ctx.Request.RequestURI, "?")
        c.Redirect(c.Ctx.Request.RequestURI[:i], 302)
        return
    }

    session, err := mgo.Dial("localhost:12345")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    // session.SetMode(mgo.Monotonic, true)

    // eventXML := XML_unmarshal(c.Lang, "event")

    coll := session.DB("juna").C("uv")

    result := UV{}
    result_pv := UV{}
    err = coll.Find(bson.M{"name": "uv_controller"}).One(&result)
    if err != nil {
        result.Name = "uv_controller"
        result.Count = 0
        err = coll.Insert(UV{"uv_controller", 0})
        if err != nil {
            log.Fatal(err)
        }
    }
    err = coll.Find(bson.M{"name": "pv_controller"}).One(&result_pv)
    if err != nil {
        result_pv.Name = "pv_controller"
        result_pv.Count = 0
        err = coll.Insert(UV{"pv_controller", 0})
        if err != nil {
            log.Fatal(err)
        }
    }

    pv := c.Controller.GetSession("pv_controller")
    uv := c.Controller.GetSession("uv_controller")
    if pv == nil {
        c.Controller.SetSession("pv_controller", int(1))
    } else {
        c.Controller.SetSession("pv_controller", pv.(int)+1)
    }
    result_pv.Count += 1

    if uv == nil {
        c.Controller.SetSession("uv_controller", int(1))
        result.Count += 1
        usrdb := session.DB("juna").C("visitor")
        userinfo := Userinfo{time.Now().UnixNano(), time.Now().String(), c.Ctx.Input.IP(), c.Lang, c.Ctx.Request.UserAgent()}
        err = usrdb.Insert(userinfo)
        if err != nil {
            log.Fatal(err)
        }
    }

    c.Data["Visit"] = result.Count
    c.Data["PageView"] = result_pv.Count

    coll.Update(bson.M{"name": "uv_controller"}, result)
    coll.Update(bson.M{"name": "pv_controller"}, result_pv)
    err = coll.Find(bson.M{"name": "uv_controller"}).One(&result)
    err = coll.Find(bson.M{"name": "pv_controller"}).One(&result_pv)

    // page start time
    // this.Data["PageStartTime"] = time.Now()

    if app, ok := c.AppController.(NestPreparer); ok {
        app.NestPrepare()
    }

}

// setLangVer sets site language version.
func (this *baseRouter) setLangVer() bool {

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

    isNeedRedir := false
    hasCookie := false

    // 1. Check URL arguments.
    lang := this.Input().Get("lang")

    // 2. Get language information from cookies.
    if len(lang) == 0 {
        lang = this.Ctx.GetCookie("lang")
        hasCookie = true
    } else {
        isNeedRedir = true
    }

    // Check again in case someone modify by purpose.
    if !i18n.IsExist(lang) {
        lang = ""
        isNeedRedir = false
        hasCookie = false
    }

    // 3. Get language information from 'Accept-Language'.
    if len(lang) == 0 {
        al := this.Ctx.Request.Header.Get("Accept-Language")
        if len(al) > 4 {
            al = al[:5] // Only compare first 5 letters.
            if i18n.IsExist(al) {
                lang = al
            }
        }
    }

    // 4. Default language is English.
    if len(lang) == 0 {
        lang = "en-US"
        isNeedRedir = false
    }

    curLang := langType{
        Lang: lang,
    }

    // Save language information in cookies.
    if !hasCookie {
        this.Ctx.SetCookie("lang", curLang.Lang, 1<<31-1, "/")
    }

    restLangs := make([]*langType, 0, len(langTypes)-1)
    for _, v := range langTypes {
        if lang != v.Lang {
            restLangs = append(restLangs, v)
        } else {
            curLang.Name = v.Name
        }
    }

    // Set language properties.
    this.Lang = lang
    this.Data["Lang"] = curLang.Lang
    this.Data["CurLang"] = curLang.Name
    this.Data["RestLangs"] = restLangs

    return isNeedRedir
}
