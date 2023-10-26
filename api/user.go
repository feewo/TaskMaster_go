package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"taskmaster/engine"
	"taskmaster/entity"
	"taskmaster/storage"
)

func (a *Api) UserCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.User
	err := decoder.Decode(&item)

	// проверка на ошибки
	if err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.UserCreate(item))
	// вывод простого текста
	// fmt.Println("ItemCreate")
	// ctx.Response.Write([]byte("ItemCreate"))
}

func (a *Api) Users(ctx *engine.Context) {
	ctx.Print((storage.UserGetAll()))
	// fmt.Println("ItemCreate")
	// ctx.Response.Write([]byte("ItemCreate"))
}

func (a *Api) User(ctx *engine.Context) {
	path := strings.Split(ctx.Request.URL.String(), "/")
	id := path[len(path)-1]

	iid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	fmt.Println(uint32(iid))
	storage.UserGet(uint32(iid))
}

func (a *Api) UserDelete(ctx *engine.Context, id string) {
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Ошибка при образовании строки в int", err)
	}
	idUint32 := uint32(idUint64)
	storage.UserDelete(idUint32)
}

func (a *Api) UserUpdate(ctx *engine.Context, id string) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.User
	err := decoder.Decode(&item)
	if err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Ошибка при образовании строки в int", err)
	}
	idUint32 := uint32(idUint64)
	ctx.Print(storage.UserUpdate(item, idUint32))
}
