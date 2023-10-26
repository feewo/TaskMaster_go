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

// ------------------ USER ------------------

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

// ------------------ TASK ------------------

func (a *Api) TaskCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.Task
	err := decoder.Decode(&item)
	// проверка на ошибки
	if err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.TaskCreate(item))
}

func (a *Api) Tasks(ctx *engine.Context) {
	ctx.Print(storage.TaskGetAll())
}

func (a *Api) Task(ctx *engine.Context) {
	path := strings.Split(ctx.Request.URL.String(), "/")
	id := path[len(path)-1]
	iid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	ctx.Print(storage.TaskGet(uint32(iid)))
}

func (a *Api) TaskDelete(ctx *engine.Context, id string) {
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Ошибка при образовании строки в int", err)
	}
	idUint32 := uint32(idUint64)
	storage.TaskDelete(idUint32)
}

func (a *Api) TaskUpdate(ctx *engine.Context, id string) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.Task
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
	ctx.Print(storage.TaskUpdate(item, idUint32))
}
