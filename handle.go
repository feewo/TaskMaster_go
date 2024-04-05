package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"taskmaster/api"
	"taskmaster/config"
	"taskmaster/db"
	"taskmaster/engine"
	"taskmaster/entity"
	"time"
)

var types map[string]bool
var hdl *api.Api
var apiMap map[string]map[string]reflect.Value

var ignoreList = make([]string, 0)

func init() {
	ignoreList = append(ignoreList, "token")
	apiMap = make(map[string]map[string]reflect.Value)
	cfg := config.Get()
	maps := cfg.Api
	hdl = &api.Api{}
	types = make(map[string]bool)
	types["ico"] = true
	types["html"] = true
	types["js"] = true
	types["svg"] = true
	types["png"] = true

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
	ctx := engine.Context{
		Response: w,
		Request:  r,
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	url := r.URL
	path := url.Path[1:]
	pathArr := strings.Split(path, "/")

	if pathArr[0] == "" || pathArr[0] == "index.html" {
		http.ServeFile(w, r, "./front/index.html")
	}
	if pathArr[0] == "css" {
		http.ServeFile(w, r, "./front/css/style.css")
	}
	if pathArr[0] == "js" {
		http.ServeFile(w, r, "./front/js/app.js")
	}
	maps, ok := apiMap[r.Method]
	if !ok {
		w.Write([]byte("Нет метода"))
		return
	}
	pathName := pathArr[0]
	if len(pathArr) > 1 {
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
		if pathName == "user" && r.Method == "POST" {
			isIgnore = true
		}
		if !isIgnore {
			token, ok := ctx.Request.Header["Authorization"]
			if !ok {
				ctx.Error(401, "Bad")
				return
			}
			var tokenDb entity.Token
			var userDb entity.User

			db.DB().Table(tokenDb.TableName()).Where("token = ? and expired > ?", token, time.Now()).Find(&tokenDb)
			fmt.Println(userDb.Role)
			if tokenDb.UserID == 0 {
				ctx.Error(401, "Bad")
				return
			}
			db.DB().Table(userDb.TableName()).Where("ID = ?", tokenDb.UserID).Find(&userDb)
			if (r.Method == "POST" || r.Method == "DELETE" || r.Method == "PUT") && userDb.Role != "admin" {
				ctx.Error(403, "Forbidden")
				return
			}
		}
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(&ctx)
		fun.Call(in)
		return
	}
}

// эта функция работает плохо
func sendFile(file string, ctx engine.Context) {
	// pwd, _ := os.Getwd()
	fmt.Println(file)
	// проблема в этой команде, она просто не работает с:
	http.ServeFile(ctx.Response, ctx.Request, file)
}

// эта функция работает хорошо
func checkStatic(path string) (string, bool) {
	// fmt.Println(path)
	typeFile := strings.Split(path, ".")
	if len(typeFile) < 2 {
		return "", false
	}
	// fmt.Println(typeFile, types[typeFile[1]])
	if _, ok := types[typeFile[1]]; ok {
		switch typeFile[1] {
		case "html":
			return "../front/" + path, true
		case "css":
			return "front/css/" + path, true
		case "js":
			return "front/js/" + path, true
		case "png", "ico", "svg":
			return "front/img/" + path, true
		}
	}
	return "", false
}
