package main

import (
	"log"
	"net/http"
)

func main() {
	eventBusRegister()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":9090", router))
}
