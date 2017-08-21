package main

import (
	"log"
	"net/http"
//	"fmt"
	"os"
    	"github.com/UnotechSoftware/hydra/v1/routes"
//    	"github.com/UnotechSoftware/hydra/v1/utils"
)

func main() {

	router := routes.NewRouter()

	var AccessLog string = "/var/log/access_log"
	var _,err = os.Stat(AccessLog)
	if os.IsNotExist(err){
	
		filep, err := os.Create(AccessLog)
		if err != nil{
		       return
		   }
		defer filep.Close()
	}
	
	ErrorLog, err := os.OpenFile("/var/log/error_log", os.O_WRONLY|os.O_CREATE, 0666)
	
	log.SetOutput(ErrorLog)
	//log.Fatal(http.ListenAndServeTLS(":443", "../server.crt", "../server.key", router))

	 log.Fatal(http.ListenAndServe(":8080", router))
}
