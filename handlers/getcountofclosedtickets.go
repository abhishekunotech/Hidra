package handlers

import (
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io"
	"io/ioutil"
	"net/http"
)

func callCountClosedTickets(w http.ResponseWriter, r *http.Request, username string, password string, customerid string) {

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getcountclosedtickets").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?CustomerID=" + customerid + "&SessionID=" + sessionIDString + "&State=close"
	client := &http.Client{}
	var bodyReader io.Reader
	req, err := http.NewRequest("GET", url, bodyReader)

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("\n\nThis caused the following error \n\n")
		logger.Error(err.Error())
	}
	req.Close = true
	bodyText, err := ioutil.ReadAll(resp.Body)
	var data interface{}
	err = json.Unmarshal(bodyText, &data)
	if err != nil {
		logger.Error(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

//Function to get the details about ticket.

func (h *Handler) GetCountofClosedTickets(w http.ResponseWriter, r *http.Request) {
	//body, _ := ioutil.ReadAll(r.Body)
	mapHttp := r.URL.Query()
	var customerid string
	var username string
	var password string
	for key, value := range mapHttp {
		if key == "CustomerID" {
			for _, valueStrg := range value {
				customerid = valueStrg
			}
		}
		if key == "UserLogin" {
			for _, valueStrg := range value {
				username = valueStrg
			}
		}
		if key == "Password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
	}
	callCountClosedTickets(w, r, username, password, customerid)

}
