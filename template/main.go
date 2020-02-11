package main

import (
	"net/http"
	"text/template"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

const tmplPath = "tmpl/"

var tmpl = template.Must(template.ParseFiles(tmplPath + "layout.html"))

func main() {
	data := TodoPageData{
		PageTitle: "My TODO List",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
