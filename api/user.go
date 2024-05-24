package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"taskmaster/context"
	"taskmaster/db"
	"taskmaster/entity"
	"taskmaster/storage"
	"time"
)

func IsAuth(ctx *context.Context, method string) {
	// получаем токен
	token, ok := ctx.Request.Header["Authorization"]
	if !ok {
		ctx.Error(401, "Bad")
		return
	}
	var tokenDb entity.Token
	// проверка токена
	db.DB().Table(tokenDb.TableName()).Where("token = ? and expired > ?", token, time.Now()).Find(&tokenDb)
	if tokenDb.UserID == 0 {
		ctx.Error(401, "Bad")
		return
	}
}

func (a *Api) UserCreate(ctx *context.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.User
	err := decoder.Decode(&item)
	// проверка на ошибки
	if err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	if item.Role == "" {
		item.Role = "user"
	}
	ctx.Print(storage.UserCreate(item))
}

func (a *Api) Users(ctx *context.Context) {
	ctx.Print(storage.UserGetAll())
}

func (a *Api) User(ctx *context.Context) {
	path := strings.Split(ctx.Request.URL.String(), "/")
	id := path[len(path)-1]
	iid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	ctx.Print(storage.UserGet(uint(iid)))
}

func (a *Api) UserToken(ctx *context.Context) {
	// получаем токен
	token, ok := ctx.Request.Header["Authorization"]
	if !ok {
		ctx.Error(401, "Bad")
		return
	}
	var tokenDb entity.Token
	// проверка токена
	db.DB().Table(tokenDb.TableName()).Where("token = ? and expired > ?", token, time.Now()).Find(&tokenDb)
	if tokenDb.UserID == 0 {
		ctx.Error(401, "Bad")
		return
	}
	ctx.Print(storage.UserGet(uint(tokenDb.UserID)))
}

func (a *Api) UserDelete(ctx *context.Context) {
	path := ctx.Request.URL.Path[1:]
	pathArr := strings.Split(path, "/")
	id := pathArr[len(pathArr)-1]
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Ошибка при образовании строки в int", err)
	}
	idUint := uint(idUint64)
	storage.UserDelete(idUint)
}

func (a *Api) UserUpdate(ctx *context.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.User
	err := decoder.Decode(&item)
	if err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	path := ctx.Request.URL.Path[1:]
	pathArr := strings.Split(path, "/")
	id := pathArr[len(pathArr)-1]
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Ошибка при образовании строки в int", err)
		ctx.Error(422, "Unprocessable Entity")
	}
	idUint := uint(idUint64)
	ctx.Print(storage.UserUpdate(item, idUint))
}

func (a *Api) UserAuth(ctx *context.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.User
	err := decoder.Decode(&item)

	// проверка на ошибки
	if err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	// вывод item
	var usr entity.User
	db.DB().Table(usr.TableName()).Where("login = ? AND password = ?", item.Login, item.Password).Find(&usr)
	if usr.ID == 0 {
		ctx.Error(401, "Unauthorized")
		return
	}

	ctx.Print(storage.UserAuth(usr))
}
func (a *Api) UserAuthDelete(ctx *context.Context) {
	token, ok := ctx.Request.Header["Authorization"]
	if !ok {
		ctx.Error(401, "Bad")
		return
	}
	storage.UserAuthDelete(token)
}
