package utils


import(
	"io/ioutil"
	"time"
	"net/http"
	"io"
	"github.com/antigloss/go/logger"
)

// Function that makes an HTTP Client execute a GET Request to the URL mentioned in the Parameter
//
// Business Logic: Will make a HTTP GET Request and return the byte converted response obtained from that request
func MakeHTTPGetCall(url string) []uint8{
	var netClient = &http.Client{
  Timeout: time.Second * 10,
}
	var bodyTess io.Reader
	newRequest,err := http.NewRequest("GET",url,bodyTess)
	start := time.Now()
	res, err := netClient.Do(newRequest)
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
