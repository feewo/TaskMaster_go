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

func (a *Api) TaskCreate(ctx *context.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.Task
	err := decoder.Decode(&item)
	if err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.TaskCreate(item))
}

func (a *Api) Tasks(ctx *context.Context) {
	ctx.Print(storage.TaskGetAll())
}

func (a *Api) TaskUser(ctx *context.Context) {
	path := strings.Split(ctx.Request.URL.String(), "/")
	id := path[len(path)-1]
	iid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	ctx.Print(storage.TaskUserGetAll(uint(iid)))
}

func (a *Api) Task(ctx *context.Context) {
	path := strings.Split(ctx.Request.URL.String(), "/")
	id := path[len(path)-1]
	iid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	ctx.Print(storage.TaskGet(uint(iid)))
}

func (a *Api) TaskDelete(ctx *context.Context) {
	path := ctx.Request.URL.Path[1:]
	pathArr := strings.Split(path, "/")
	idString := pathArr[1]
	// преобразование в uint
	idUint64, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		fmt.Println("Ошибка при образовании строки в int", err)
		return
	}
	id := uint(idUint64)
	storage.TaskDelete(id)
}

func (a *Api) TaskUpdate(ctx *context.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	// var taskMap map[string]interface{}
	var task entity.Task
	err := decoder.Decode(&task)
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
	ctx.Print(storage.TaskUpdate(task, idUint))
}
