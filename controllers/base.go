package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"wechat/g"
)

type Checker interface {
	CheckLogin()
}

type BaseController struct {
	beego.Controller
	IsAdmin bool
}

func (this *BaseController) Prepare() {
	this.AssignIsAdmin()
	if app, ok := this.AppController.(Checker); ok {
		app.CheckLogin()
	}
}

func (this *BaseController) AssignIsAdmin() {
	bb_name := this.Ctx.GetCookie("wechat_name")
	bb_password := this.Ctx.GetCookie("wechat_password")
	if bb_name == "" || bb_password == "" {
		this.IsAdmin = false
		return
	}

	if bb_name != g.RootName || bb_password != g.RootPass {
		this.IsAdmin = false
	}

	this.IsAdmin = true
	this.Data["IsAdmin"] = this.IsAdmin
}

func (this *BaseController) GetIntWithDefault(paramKey string, defaultVal int) int {
	valStr := this.GetString(paramKey)
	var val int
	if valStr == "" {
		val = defaultVal
	} else {
		var err error
		val, err = strconv.Atoi(valStr)
		if err != nil {
			val = defaultVal
		}
	}
	return val
}

func (this *BaseController) JsStorage(action, key string, values ...string) {
	value := action + ":::" + key
	if len(values) > 0 {
		value += ":::" + values[0]
	}
	this.Ctx.SetCookie("JsStorage", value, 1<<31-1, "/")
}
