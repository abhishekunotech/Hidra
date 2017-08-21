//This is the entry point for a middleware to Felicity Components
//Install using go install in this directory
//Running this will start up an HTTPS Server instance on port 443, and can be accessed at http://localhost:443.
//You will first need to generate SSL Certificates for the server using instructions from README.md, if you do not already have any certificates
//Author: Operations Management Team - Unotech Software

package main
import (
   "github.com/abhishekunotech/hydra/v1/routes"
   "log"
   "net/http"
)

func main() {

//Implements a router.  The router defines the list of all available APIs  

   router := routes.NewRouter()

//ListenAndServeTLS starts an HTTPS server.
//Change the first and second parameters as per the locations of your certificates

   go log.Fatal(http.ListenAndServeTLS(":443", "../server.crt", "../server.key", router))
}











