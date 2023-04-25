package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Welcome to my website. You've requested: %s\n", request.URL.Path)
	})

	r.HandleFunc("/books/{title}/page/{page}", func(response http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(response, "You've requested the book: %s on page %s\n", title, page)
	})

	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3000", nil)
}
