package dto

import (
	"encoding/json"
	"taskmaster/db"
	"taskmaster/engine"
	"taskmaster/entity"
)

type Api struct{}

func (a *Api) UserAuth(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.User
	err := decoder.Decode(&item)
	if err != nil {
		ctx.Error(400, "Bad")
	}
	var usr entity.User
	db.DB().Table(usr.TableName()).Where("login = ? and password = ?",
		usr.Login, usr.Password).Find(&usr)
	if usr.ID == 0 {
		ctx.Error(401, "Bad")
		return
	}
}
