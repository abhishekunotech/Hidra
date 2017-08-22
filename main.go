//This is the entry point for a middleware to Felicity Components
//Install using go install in this directory
//Running this will start up an HTTPS Server instance on port 443, and can be accessed at http://localhost:443.
//You will first need to generate SSL Certificates for the server using instructions from README.md, if you do not already have any certificates
//Author: Operations Management Team - Unotech Software

package main
import (

	"log"
	"net/http"
//	"fmt"
	"os"
	"github.com/antigloss/go/logger"
    	"github.com/Unotechsoftware/Hydra/routes"
//    	"github.com/UnotechSoftware/hydra/v1/utils"


)

func main() {

//Implements a router.  The router defines the list of all available APIs  

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
  
  //ListenAndServeTLS starts an HTTPS server.
//Change the first and second parameters as per the locations of your certificates

	 log.Fatal(http.ListenAndServe(":8080", router))

}










