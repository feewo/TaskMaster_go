package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"taskmaster/context"
	"taskmaster/entity"
	"taskmaster/storage"
)

func (a *Api) TaskPointCreate(ctx *context.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.TaskPoint
	err := decoder.Decode(&item)
	if err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.TaskPointCreate(item))
}

func (a *Api) TaskPoints(ctx *context.Context) {
	ctx.Print(storage.TaskPointGetAll())
}

func (a *Api) TaskPoint(ctx *context.Context) {
	path := strings.Split(ctx.Request.URL.String(), "/")
	id := path[len(path)-1]
	iid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	ctx.Print(storage.TaskPointGet(uint(iid)))
}

func (a *Api) TaskPointTask(ctx *context.Context) {
	path := strings.Split(ctx.Request.URL.String(), "/")
	id := path[len(path)-1]
	iid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	ctx.Print(storage.TaskPointTaskGet(uint(iid)))
}

func (a *Api) TaskPointDelete(ctx *context.Context) {
	path := ctx.Request.URL.Path[1:]
	pathArr := strings.Split(path, "/")
	id := pathArr[len(pathArr)-1]
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Ошибка при образовании строки в int", err)
	}
	idUint := uint(idUint64)
	storage.TaskPointDelete(idUint)
}

func (a *Api) TaskPointUpdate(ctx *context.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var taskPoint entity.TaskPoint
	err := decoder.Decode(&taskPoint)
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
	}
	idUint := uint(idUint64)
	ctx.Print(storage.TaskPointUpdate(taskPoint, idUint))
}
