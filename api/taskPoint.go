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

func (a *Api) TaskPointCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.TaskPoint
	err := decoder.Decode(&item)
	// проверка на ошибки
	if err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.TaskPointCreate(item))
}

func (a *Api) TaskPoints(ctx *engine.Context) {
	ctx.Print(storage.TaskPointGetAll())
}

func (a *Api) TaskPoint(ctx *engine.Context) {
	path := strings.Split(ctx.Request.URL.String(), "/")
	id := path[len(path)-1]
	iid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	ctx.Print(storage.TaskPointGet(uint(iid)))
}

func (a *Api) TaskPointDelete(ctx *engine.Context) {
	path := ctx.Request.URL.Path[1:]
	pathArr := strings.Split(path, "/")
	id := pathArr[1]
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Ошибка при образовании строки в int", err)
	}
	idUint := uint(idUint64)
	storage.TaskPointDelete(idUint)
}

func (a *Api) TaskPointUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	// var taskPointMap map[string]interface{}
	var taskPoint entity.TaskPoint
	err := decoder.Decode(&taskPoint)
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
	idUint := uint(idUint64)
	// ctx.Print(storage.TaskPointUpdate(taskPointMap, idUint))
	ctx.Print(storage.TaskPointUpdate(taskPoint, idUint))
}
