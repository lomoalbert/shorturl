package controllers

import (
	"github.com/lomoalbert/shorturl/models"

	"github.com/astaxie/beego"
	"net/url"
)

// Operations about object
type ShortURLController struct {
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