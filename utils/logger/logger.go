package logger

import (
	"github.com/antigloss/go/logger"	
    "fmt"
   // "log"
    "net/http"
    "time"
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

	logger.Info(string(s))
	
    })
}
