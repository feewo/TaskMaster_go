package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"taskmaster/api"
	"taskmaster/config"
	"taskmaster/db"
	"taskmaster/engine"
	"taskmaster/entity"
<<<<<<< Updated upstream
	_ "taskmaster/entity"
=======
>>>>>>> Stashed changes
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
	url := r.URL
	path := url.Path[1:]
	pathArr := strings.Split(path, "/")
	// fmt.Println(reflect.TypeOf(path))
	// if pathArr[0] == "" {
	// 	sendFile("front/index.html", ctx)
	// 	return
	// }
	if stat, ok := checkStatic(path); ok {
		fmt.Println(stat, "ok")
		sendFile(stat, ctx)
		return
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
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
			if s == pathName {
				isIgnore = true
				break
			}
		}
		if !isIgnore {
			token, ok := ctx.Request.Header["Authentication"]
			if !ok {
				ctx.Error(401, "Bad")
				return
			}
			var tokenDb entity.Token
<<<<<<< Updated upstream
			db.DB().Table(tokenDb.TableName()).Where("token = ? and expired > ?", token, time.Now())
=======
			var userDb entity.User

			db.DB().Table(tokenDb.TableName()).Where("token = ? and expired > ?", token, time.Now()).Find(&tokenDb)

>>>>>>> Stashed changes
			if tokenDb.Iid == 0 {
				ctx.Error(401, "Bad")
				return
			}
<<<<<<< Updated upstream
=======
			db.DB().Table(userDb.TableName()).Where("iid = ?", tokenDb.Iid).Find(&userDb)
			if (r.Method == "POST" || r.Method == "DELETE" || r.Method == "PUT") && userDb.Role == "user" {
				ctx.Error(401, "Bad")
				return
			}
>>>>>>> Stashed changes
		}
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(&ctx)
		fun.Call(in)
		return
	}

	// POST
	// if r.Method == "POST" {
	// 	switch pathArr[0] {
	// 	case "user":
	// 		hdl.UserCreate(ctx)
	// 	default:
	// 		w.Write([]byte("Нет БО"))
	// 	}
	// 	return
	// }
	// // GET
	// if r.Method == "GET" {
	// 	switch pathArr[0] {
	// 	case "user":
	// 		if len(pathArr) > 1 {
	// 			id := pathArr[1]
	// 			hdl.User(&ctx, id)
	// 		} else {
	// 			hdl.Users(ctx)
	// 		}
	// 	default:
	// 		w.Write([]byte("Нет БО"))
	// 	}
	// 	return
	// }

	// // DELETE
	// if r.Method == "DELETE" {
	// 	id := pathArr[1]
	// 	switch pathArr[0] {
	// 	case "user":
	// 		if id != "" {
	// 			hdl.UserDelete(ctx, id)
	// 		} else {
	// 			w.Write([]byte("Нет БО"))
	// 		}
	// 	default:
	// 		w.Write([]byte("Нет БО"))
	// 	}
	// 	return
	// }
	// // PUT
	// if r.Method == "PUT" {
	// 	id := pathArr[1]
	// 	switch pathArr[0] {
	// 	case "user":
	// 		if id != "" {
	// 			hdl.UserUpdate(ctx, id)
	// 		} else {
	// 			w.Write([]byte("Нет БО"))
	// 		}
	// 	default:
	// 		w.Write([]byte("Нет БО"))
	// 	}
	// 	return
	// }
}
func sendFile(file string, ctx engine.Context) {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	http.ServeFile(ctx.Response, ctx.Request, pwd+"/")
}

func checkStatic(path string) (string, bool) {
	fmt.Println(path)
	typeFile := strings.Split(path, ".")
	if len(typeFile) < 2 {
		return "", false
	}
	fmt.Println(typeFile, types[typeFile[1]])
	if _, ok := types[typeFile[1]]; ok {
		switch typeFile[1] {
		case "html":
			return "front/" + path, true
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
