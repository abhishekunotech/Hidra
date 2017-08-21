package logger

import (
    "fmt"
   // "log"
    "net/http"
    "time"
    "os"
)

func Logger(inner http.Handler, name string) http.Handler {
 return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

	inner.ServeHTTP(w, r)

	var s string
	s = fmt.Sprintf("%s::\t%s\t%s\t%s\t%s\t%s\n",
            start,
	    r.Method,
	    r.RemoteAddr,
            r.RequestURI,
            name,
            time.Since(start),)

	fmt.Println(string(s))

	fp, err := os.OpenFile("/var/log/access_log", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil{
	   panic(err)
	}
	
        //data := []byte(s)
	
	_, err = fp.WriteString(s)
	if err != nil{
	   panic(err)
	}
	defer fp.Close()

    })
}
