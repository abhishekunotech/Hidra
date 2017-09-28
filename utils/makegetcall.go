package utils


import(
	"io/ioutil"
	"time"
	"net/http"
	"github.com/antigloss/go/logger"
)


func MakeHTTPGetCall(url string) []uint8{
	start := time.Now()
	res, err := http.Get(url)
        if err != nil {
                logger.Error(err.Error())
        }

	logger.Info("Time Taken to make a call to component with URL "+url+" is")
	till := time.Since(start).String()
	logger.Info(till)
	//return res
	
	bodyText, err := ioutil.ReadAll(res.Body)
	return bodyText
}
