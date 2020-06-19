package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/mjd/bee-api-gs/models"
	"strconv"
)

// Operations about Wine
type WineController struct {
	beego.Controller
}

// @Title CreateWine
// @Description Add a new wine
// @Param	body		body 	models.Wine	true		"body for user content"
// @Success 200 {int} models.Wine.Id
// @Failure 403 body is empty
// @router / [post]
func (u *WineController) Post() {
	var wine models.Wine
	json.Unmarshal(u.Ctx.Input.RequestBody, &wine)
	uu, err := models.AddWine(wine)
	if err != nil {
		// handle error
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = uu	//[string]string{"uid": string(uid)}
	}

	u.ServeJSON()
}

// @Title GetAll
// @Description get all Wines
// @Success 200 {object} models.Wine
// @router / [get]
func (u *WineController) GetAllWines() {
	wines, err := models.GetAllWines()
	if err != nil {

	}

	u.Data["json"] = wines
	u.ServeJSON()
}

// @Title Get
// @Description get wine by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Wine
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *WineController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		i, err := strconv.Atoi(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			wine, err := models.GetWine(i)
			if err != nil {
				u.Data["json"] = err.Error()
			} else {
				u.Data["json"] = wine
			}
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the wine
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Wine	true		"body for user content"
// @Success 200 {object} models.Wine
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *WineController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var wine models.Wine
		json.Unmarshal(u.Ctx.Input.RequestBody, &wine)
		uu, err := models.UpdateWine(uid, wine)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the wine
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *WineController) Delete() {
	uid := u.GetString(":uid")
	i, err := strconv.Atoi(uid)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		models.DeleteWine(i)
		u.Data["json"] = "delete success!"
		u.ServeJSON()
	}
}


