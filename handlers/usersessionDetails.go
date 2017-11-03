package handlers

import (
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io"
	"time"
	"io/ioutil"
	"net/http"
)

type CustomerSessionObject struct {
	SessionIDStrg string `json:"SessionID"`
}

func callCustomerSessionDetails(username string, password string) string {

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.CustomerSessionApi").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?CustomerUserLogin=" + username + "&Password=" + password
	client := &http.Client{}
	var bodyReader io.Reader
	start := time.Now()
	req, err := http.NewRequest("GET", url, bodyReader)
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("\n\n Session Creation failed because - \n\n")
		logger.Error(err.Error())
	}
	req.Close = true
	till := time.Since(start).String()

	logger.Info(" The call to URL "+url+" took "+till+" to execute.")
	bodyText, err := ioutil.ReadAll(resp.Body)
	var data CustomerSessionObject
	err = json.Unmarshal(bodyText, &data)
	if err != nil {
		logger.Error("\n\n Error Occured in unmarshalling Session JSON \n\n")
		logger.Error(err.Error())
	}
	return data.SessionIDStrg
}
