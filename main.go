package goson_http

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Building REST APIs with Go 1.22!")

	mux := http.NewServeMux() // New HTTP Multiplexer

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
