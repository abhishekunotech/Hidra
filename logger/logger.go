/*
	This function will wrap the passed handler with logging and timing functionality.
	Install using go install in this directory.

	Author: Operations Management Team - Unotech Software.
*/

package logger

import (
	"fmt"
	// "log"
	"net/http"
	"os"
	"time"
)


//Function logger will accept the request and start the logging. 
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//Current time will be set to the start.
		start := time.Now()
//ServeHTTP method simply checks the Host header on the request.

		inner.ServeHTTP(w, r)

		var s string
//This will display method of http request, remote address,URL of request and time required to process the request. 
		s = fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n",
			r.Method,
			r.RemoteAddr,
			r.RequestURI,
			name,
			time.Since(start))

		fmt.Println(string(s))

//This will open a file named middleware.log with write only mode (0600-before umask) and data is appended to the file.

		fp, err := os.OpenFile("/tmp/middleware.log", os.O_APPEND|os.O_WRONLY, 0600)
//Panic is a built-in function that stops the ordinary flow of control and begins panicking.
//Panics can be initiated by invoking panic directly. They can also be caused by runtime errors.
		if err != nil {
			panic(err)
		}

		//data := []byte(s)
//fp.WriteString writes the contents of the string s to fp
		n, err := fp.WriteString(s)
//Errors are handled if any.
		if err != nil {
			panic(err)
		}

//By introducing defer statement, it is ensured that the files are always closed. 
		defer fp.Close()
//Display the contents written using fp.WriteString
		fmt.Println("bytes :", n)
		/* 	log.Printf(
		            "%s\t%s\t%s\t%s\t%s",
		            r.Method,
			    r.RemoteAddr,
		            r.RequestURI,
		            name,
		            time.Since(start),
		        )*/
	})
}


