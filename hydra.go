package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	go log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", router))
}
