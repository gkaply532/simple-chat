package main

import (
	"embed"
	"net/http"
)

//go:embed static
var static embed.FS

func init() {
	http.Handle("/static/", http.FileServer(http.FS(static)))
}
