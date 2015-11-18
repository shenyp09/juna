package controllers

import (
    "encoding/xml"
    "fmt"
    "github.com/astaxie/beego"
    // "github.com/astaxie/beego/context"
    // "github.com/beego/i18n"
    "io/ioutil"
    "os"
)

func XML_unmarshal(lang string) *Recurlynews {
    file, err := os.Open("static/data/news_" + lang + ".xml") // For read access.
    if err != nil {
        fmt.Printf("error: %v", err)
        return nil
    }
    defer file.Close()
    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Printf("error: %v", err)
        return nil
    }
    v := Recurlynews{}
    err = xml.Unmarshal(data, &v)
    if err != nil {
        fmt.Printf("error: %v", err)
        return nil
    }
    return &v
}

func session_controller(c *beego.Controller) {
    sess := c.GetSession("pv_controller")
    c.SetSession("uv_controller", true)
    if sess == nil {
        c.SetSession("pv_controller", int(1))
        c.Data["Visit"] = 0
    } else {
        c.SetSession("pv_controller", sess.(int)+1)
        c.Data["Visit"] = sess.(int)
    }
}

type Recurlynews struct {
    XMLName xml.Name `xml:"newslist"`
    Version string   `xml:"version,attr"`
    Nws     []news   `xml:"news"`
    // Description string   `xml:",innerxml"`
}

type news struct {
    XMLName     xml.Name `xml:"news"`
    Title       string   `xml:"title"`
    ID          string   `xml:"id"`
    Focus       string   `xml:"focus"`
    Date        string   `xml:"date"`
    Description string   `xml:"description"`
    Content     Content  `xml:"content"`
    Image       string   `xml:"image"`
    link        string   `xml:"link"`
}

type Content struct {
    Description string `xml:",innerxml"`
}
