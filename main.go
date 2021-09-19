package main

import (
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/handlers"
)

func main() {

	http.HandleFunc("/item/", handlers.ReadCSV)
	http.ListenAndServe(":8080", nil)
}
