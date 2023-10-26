package main

import (
	"fmt"
	"net/http"
	"taskmaster/config"
	"time"
)

func main() {
	cfg := config.Get()

	// СЕРВЕР
	m := http.NewServeMux()
	m.Handle("/", http.HandlerFunc(handle)) // не работает, но выглядит примерно так :( (undefined: handle)
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      m,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Listen port" + cfg.Server.Port)
	server.ListenAndServe()
}
