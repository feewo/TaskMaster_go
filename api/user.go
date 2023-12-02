package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"taskmaster/db"
	"taskmaster/engine"
	"taskmaster/entity"
	"taskmaster/storage"
	"time"

	"github.com/google/uuid"
)

func TokenGet(usr entity.User) entity.Token {
	var lastToken entity.Token
	db.DB().Table("token").Order("tokid DESC").Last(&lastToken)

	lastTokenID := lastToken.Tokid
	token := entity.Token{
		Tokid:   lastTokenID + 1,
		Iid:     usr.Iid,
		Token:   uuid.NewString(),
		Expired: time.Now().Add(1 * time.Hour),
	}
	return token
}

func (a *Api) UserCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.User
	err := decoder.Decode(&item)
	// проверка на ошибки
	if err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	var lastUser entity.User
	db.DB().Table("user").Order("iid DESC").Last(&lastUser)
	item.Iid = lastUser.Iid + 1
	item.Role = "user"
	ctx.Print(storage.UserCreate(item))
}

func (a *Api) Users(ctx *engine.Context) {
	ctx.Print(storage.UserGetAll())
}

func (a *Api) User(ctx *engine.Context) {
	path := strings.Split(ctx.Request.URL.String(), "/")
	id := path[len(path)-1]
	iid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	ctx.Print(storage.UserGet(uint32(iid)))
}

func (a *Api) UserDelete(ctx *engine.Context) {
	path := ctx.Request.URL.Path[1:]
	pathArr := strings.Split(path, "/")
	id := pathArr[1]
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Ошибка при образовании строки в int", err)
	}
	idUint32 := uint32(idUint64)
	storage.UserDelete(idUint32)
}

func (a *Api) UserUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.User
	err := decoder.Decode(&item)
	if err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	path := ctx.Request.URL.Path[1:]
	pathArr := strings.Split(path, "/")
	id := pathArr[1]
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Ошибка при образовании строки в int", err)
	}
	idUint32 := uint32(idUint64)
	ctx.Print(storage.UserUpdate(item, idUint32))
}

func (a *Api) UserAuth(ctx *engine.Context) {
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
	db.DB().Table(usr.TableName()).Where("login = ?", item.Login, " and password = ?", item.Password).Find(&usr)
	if usr.Iid == 0 {
		ctx.Error(401, "Bad")
		return
	}

	ctx.Print(storage.UserAuth(usr))
}
