package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello, world! By Dagger-poc!/n")
	})
	http.ListenAndServe(":9090", nil)
}
