package main

import (
	"net/http"
)

func front(path string, w http.ResponseWriter, r *http.Request) {
	if path == "" {
		http.ServeFile(w, r, "./front/index.html")
	}
	http.ServeFile(w, r, "./front/"+path)
}
