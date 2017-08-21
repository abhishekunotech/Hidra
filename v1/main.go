package main

import (
	"github.com/abhishekunotech/hydra/v1/routes"
	"log"
	"net/http"
//	"fmt"
	"os"
	"github.com/antigloss/go/logger"
    	"github.com/Unotechsoftware/Hydra/v1/routes"
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
	logger.Init("./log", // specify the directory to save the logfiles
			400, // maximum logfiles allowed under the specified log directory
			20, // number of logfiles to delete when number of logfiles exceeds the configured limit
			100, // maximum size of a logfile in MB
			false) // whether logs with Trace level are written down
logger.Info("Some Info")
	 log.Fatal(http.ListenAndServe(":8080", router))
}
