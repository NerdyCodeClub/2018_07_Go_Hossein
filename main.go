package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}

func init() {
	EntitiesController.Init()
}
