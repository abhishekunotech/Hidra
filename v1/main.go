package main

import (
	"log"
	"net/http"
    "github.com/UnotechSoftware/felicitymiddleware/v1/routes"
)

func main() {

	router := routes.NewRouter()

	//log.Fatal(http.ListenAndServeTLS(":443", "../server.crt", "../server.key", router))

	go log.Fatal(http.ListenAndServe(":8080", router))
}
