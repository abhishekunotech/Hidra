package utils

import(
	"time"
	"net/http"
	"bytes"
	"io/ioutil"
	"github.com/antigloss/go/logger"
)


func MakeHTTPPostCall(url string,b *bytes.Buffer) []uint8{
	client := &http.Client{}
	start := time.Now()
        req, err := http.NewRequest("POST", url, b)

        if err != nil {
                logger.Error("\n\n Request to Create Request Failed \n\n")
                logger.Error(err.Error())
        }

        logger.Info("Request")

        req.Close = true
        req.Header.Set("Content-Type", "application/json")
        resp, err := client.Do(req)
        if err != nil {
                logger.Error("\n\n POST REQUEST TO FELICITY FAILED \n\n")
                logger.Error(err.Error())
        }
        //req.Close = true
	till := time.Since(start).String()

	logger.Info("The URL "+url+" took "+till+" to execute")
        bodyText, err := ioutil.ReadAll(resp.Body)
	
	if err != nil{
		logger.Error(err.Error())
	}
	return bodyText
}
