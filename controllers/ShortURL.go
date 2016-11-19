package controllers

import (
	"github.com/lomoalbert/shorturl/models"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/astaxie/beego"
	"net/url"
	"strconv"
)

// Operations about object
type ShortURLController struct {
	beego.Controller
}

type RedirectController struct {
	beego.Controller
}

type HomeController struct {
	beego.Controller
}

// @Title Create
// @Description create ShortURL
// @Param	url		formData 	string	true		"The full URL"
// @Success 200 {object} models.SURL
// @Failure 403 body is empty
// @router / [post]
func (ctl *ShortURLController) Post() {
	longurl := ctl.GetString("url")
	_,err :=url.Parse(longurl)
	if err != nil{
		ctl.Abort("400")
		return
	}
	surl := &models.SURL{LongURL:longurl}
	surl.Save()
	beego.Debug(surl)
	ctl.Data["json"]=surl
	ctl.ServeJSON()
}


// @Title Get
// @Description get ShortURL
// @Param	code		path 	string	true		"The short code"
// @Success 200 {object} models.SURL
// @Failure 403 body is empty
// @router /:code [get]
func (ctl *ShortURLController) Get() {
	code := ctl.GetString(":code")
	surl := &models.SURL{ShortCode:code}
	surl.Get()
	beego.Debug(surl)
	ctl.Data["json"]=surl
	ctl.ServeJSON()
}

// @Title Get QrCode Image
// @Description get QRcode Image
// @Param	code		path 	string	true		"The short code"
// @Success 200 {object} models.SURL
// @Failure 403 body is empty
// @router /:code/qr [get]
func (ctl *ShortURLController) Qr() {
	code := ctl.GetString(":code")
	data,err := qrcode.Encode(ctl.Ctx.Input.Scheme()+"://"+ctl.Ctx.Input.Host()+":"+strconv.Itoa(ctl.Ctx.Input.Port())+"/"+code,qrcode.Medium,256)
	if err!=nil{
		ctl.Abort("500")
	}
	ctl.Ctx.ResponseWriter.Write(data)
}

func (ctl *RedirectController) Get(){
	code := ctl.GetString(":code")
	surl := &models.SURL{ShortCode:code}
	surl.Get()
	ctl.Redirect(surl.LongURL ,301)
}

func (ctl *HomeController)Get(){
	beego.Debug("HomeController")
	ctl.TplName = "index.html"
}