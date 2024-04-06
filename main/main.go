package main

import (
	"fmt"
	"net/http"
	"taskmaster/config"
	"taskmaster/db"
	"time"
)

func main() {
	cfg := config.Get()

	// СЕРВЕР
	m := http.NewServeMux()
	m.Handle("/", http.HandlerFunc(handle))
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      m,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	db.Migrate()
	fmt.Println("Listen port" + cfg.Server.Port)
	server.ListenAndServe()
}
