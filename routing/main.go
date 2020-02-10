package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
)

func main() {
	g := gorillaRouter()
	c := chiRouter()
	go http.ListenAndServe(":80", g)
	http.ListenAndServe(":8080", c)
}

func gorillaRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	return r
}

func chiRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		title := chi.URLParam(r, "title")
		page := chi.URLParam(r, "page")

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})
	return r
}
