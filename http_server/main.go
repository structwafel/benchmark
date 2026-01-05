package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}
