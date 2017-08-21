package main

import (
	"github.com/abhishekunotech/hydra/v1/routes"
	"log"
	"net/http"
)

func main() {

	router := routes.NewRouter()

	go log.Fatal(http.ListenAndServeTLS(":443", "../server.crt", "../server.key", router))
}
