package utils


import(
	"time"
	"net/http"
	"github.com/antigloss/go/logger"
)


func MakeHTTPGetCall(url string) *http.Response{
	start := time.Now()
	res, err := http.Get(url)
        if err != nil {
                logger.Error(err.Error())
        }

	logger.Info("Time Taken to make a call to component with URL "+url+" is")
	till := time.Since(start).String()
	logger.Info(till)
	return res
}
