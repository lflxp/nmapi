package controllers

import (
	"encoding/json"

	"github.com/lflxp/nmapi/models"
)

type NmapController struct {
	BaseController
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Scanner	true		"The object content"
// @Success 200 {object} models.Scanner
// @Failure 403 body is empty
// @router / [post]
func (this *NmapController) Post() {
	var ob models.Scanner
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		data, err := ob.Parse()
		if err != nil {
			this.Data["json"] = err.Error()
		} else {
			this.Data["json"] = data
		}
	}

	this.ServeJSON()
}
