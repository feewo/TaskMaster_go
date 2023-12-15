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
)

func (a *Api) TaskCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var item entity.Task
	err := decoder.Decode(&item)
	// проверка на ошибки
	if err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	var lastTask entity.Task
	db.DB().Table("task").Order("tid DESC").Last(&lastTask)
	item.Tid = lastTask.Tid + 1
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

func (a *Api) TaskDelete(ctx *engine.Context) {
	path := ctx.Request.URL.Path[1:]
	pathArr := strings.Split(path, "/")
	id := pathArr[1]
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Ошибка при образовании строки в int", err)
	}
	idUint32 := uint32(idUint64)
	storage.TaskDelete(idUint32)
}

func (a *Api) TaskUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var taskMap map[string]interface{}
	err := decoder.Decode(&taskMap)
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
	ctx.Print(storage.TaskUpdate(taskMap, idUint32))
}
