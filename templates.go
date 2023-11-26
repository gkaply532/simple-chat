package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed templates
var templatesFS embed.FS
var templates = must(template.New("").Funcs(template.FuncMap{
	"safeHTML": func(html string) template.HTML {
		return template.HTML(html)
	},
}).ParseFS(templatesFS, "templates/*"))

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
}
