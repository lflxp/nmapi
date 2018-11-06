package controllers

/*
	1.允许Headers表头通过验证 YC-Token
	2.增加Options 功能验证
	3.允许服务器跨域访问
*/
import (
	"encoding/base64"
	"strings"

	"github.com/astaxie/beego"
	"github.com/lflxp/nmapi/pkg"
)

type Other struct {
	User  string
	Pwd   string
	Name  string
	Token string
}

type BaseController struct {
	beego.Controller
	Others Other
}

// 改为正则表达式
func IsWhiteList(name string) bool {
	rs := false
	white_list := []string{
		"/cdl/api/v1.0/token",
		// "/cdl/api/v1.0/helm/search",
	}
	for _, x := range white_list {
		if name == x {
			rs = true
			break
		}
	}
	return rs
}

func (this *BaseController) Prepare() {
	beego.ReadFromRequest(&this.Controller)
	// token vaild
	if IsWhiteList(this.Ctx.Input.URL()) {
		this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")                                   //允许访问源
		this.Ctx.Output.Header("Access-Control-Allow-Methods", "DELETE,POST, GET, PUT, OPTIONS")     //允许post访问
		this.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Token") //header的类型
		this.Ctx.Output.Header("Access-Control-Max-Age", "1728000")
		this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
		this.Ctx.Output.Header("content-type", "application/json") //返回数据格式是json
	} else {
		tmp := strings.Split(this.Ctx.Input.Header("Authorization"), " ")
		if len(tmp) == 1 {
			// beego.Critical("1", tmp)
			this.Data["json"] = "Permission Denied"
			this.ServeJSON()
		} else if len(tmp) > 1 {
			beego.Critical("2", tmp, len(tmp))
			decodeBytes, _ := base64.StdEncoding.DecodeString(tmp[1])
			this.Others.Name = string(decodeBytes)
			this.Others.Token = string(decodeBytes)[:len(string(decodeBytes))-1]
			t := pkg.GetToken(this.Others.Token)
			if t == nil {
				this.Data["json"] = "token is useless"
				this.ServeJSON()
			} else {
				if t.IsExpire() {
					tmps := strings.Split(t.Username, ":")
					this.Others.User = tmps[0]
					this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")                                   //允许访问源
					this.Ctx.Output.Header("Access-Control-Allow-Methods", "DELETE,POST, GET, PUT, OPTIONS")     //允许post访问
					this.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Token") //header的类型
					this.Ctx.Output.Header("Access-Control-Max-Age", "1728000")
					this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
					this.Ctx.Output.Header("content-type", "application/json") //返回数据格式是json
				} else {
					beego.Critical(t)
					this.Data["json"] = "token is expire"
					this.ServeJSON()
				}
			}
		}
	}

	// this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")                                   //允许访问源
	// this.Ctx.Output.Header("Access-Control-Allow-Methods", "DELETE,POST, GET, PUT, OPTIONS")     //允许post访问
	// this.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Token") //header的类型
	// this.Ctx.Output.Header("Access-Control-Max-Age", "1728000")
	// this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	// this.Ctx.Output.Header("content-type", "application/json") //返回数据格式是json
}

func (this *BaseController) AllowCross() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")                                   //允许访问源
	this.Ctx.Output.Header("Access-Control-Allow-Methods", "DELETE,POST, GET, PUT, OPTIONS")     //允许post访问
	this.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Token") //header的类型
	this.Ctx.Output.Header("Access-Control-Max-Age", "1728000")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	this.Ctx.Output.Header("content-type", "application/json") //返回数据格式是json
}

func (this *BaseController) Options() {
	this.AllowCross() //允许跨域
	this.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	this.ServeJSON()
}
