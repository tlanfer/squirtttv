package ui

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed www/dist/**
var www embed.FS

func NewHandler() http.Handler {
	sub, _ := fs.Sub(www, "www/dist")
	fileServer := http.FileServer(http.FS(sub))
	return serveIndex(fileServer)
}

func serveIndex(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/assets") {
			r.URL.Path = "/"
		}
		next.ServeHTTP(w, r)
	})
}
