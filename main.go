package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("HELLO")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello le monde"))
	})
	http.ListenAndServe(":3000", nil)
}
