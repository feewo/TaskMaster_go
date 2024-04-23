package main

import (
	"net/http"
	"reflect"
	"strings"
	"taskmaster/api"
	"taskmaster/config"
	"taskmaster/context"
)

var hdl *api.Api
var apiMap map[string]map[string]reflect.Value

var ignoreList = make([]string, 0)

func init() {
	ignoreList = append(ignoreList, "token")
	apiMap = make(map[string]map[string]reflect.Value)
	cfg := config.Get()
	maps := cfg.Api
	hdl = &api.Api{}

	services := reflect.ValueOf(hdl)
	_struct := reflect.TypeOf(hdl)
	for methodNum := 0; methodNum < _struct.NumMethod(); methodNum++ {
		method := _struct.Method(methodNum)
		val, ok := maps[method.Name]
		if !ok {
			continue
		}
		if _, ok := apiMap[val.Method]; !ok {
			apiMap[val.Method] = make(map[string]reflect.Value)
		}
		apiMap[val.Method][val.Url] = services.Method(methodNum)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.Context{
		Response: w,
		Request:  r,
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	url := r.URL
	path := url.Path[1:]
	pathArr := strings.Split(path, "/")
	if pathArr[0] != "api" {
		front(path, w, r)
		return
	}
	maps, ok := apiMap[r.Method]
	if !ok {
		w.Write([]byte("Нет метода"))
		return
	}
	pathName := pathArr[1]
	if len(pathArr) > 2 {
		pathName += "/{id}"
	}

	if fun, ok := maps[pathName]; ok {
		isIgnore := false
		for _, s := range ignoreList {
			if s == pathName {
				isIgnore = true
				break
			}
		}
		// регстрация не требует авторизации
		if pathName == "user" && r.Method == "POST" {
			isIgnore = true
		}
		// проверяем токен и роли
		if !isIgnore {
			api.IsAuth(&ctx, r.Method)
		}
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(&ctx)
		fun.Call(in)
		return
	}
}
