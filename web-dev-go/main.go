package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Hello, you've requested: %s\n", request.URL.Path)
	})

	http.ListenAndServe(":3000", nil)
}
