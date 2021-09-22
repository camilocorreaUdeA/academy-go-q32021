package main

import (
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/handlers"
)

const (
	itemsRoute = "/items/"
)

func main() {

	http.HandleFunc(itemsRoute, handlers.GetItems)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("An error ocurred trying to run the service: %s", err.Error())
	}
}
