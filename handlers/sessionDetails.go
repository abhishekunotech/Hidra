package handlers

import (
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io"
	"io/ioutil"
	"net/http"
)

type SessionObject struct {
	SessionIDStrg string `json:"SessionID"`
}

func callSessionDetails(username string, password string) string {

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.SessionAPI").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password
	client := &http.Client{}
	var bodyReader io.Reader
	req, err := http.NewRequest("GET", url, bodyReader)
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("\n\n Session Creation failed because - \n\n")
		logger.Error(err.Error())
	}
	req.Close = true
	bodyText, err := ioutil.ReadAll(resp.Body)
	var data SessionObject
	err = json.Unmarshal(bodyText, &data)
	if err != nil {
		logger.Error("\n\n Error Occured in unmarshalling Session JSON \n\n")
		logger.Error(err.Error())
	}
	return data.SessionIDStrg
}
