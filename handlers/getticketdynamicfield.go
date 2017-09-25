
package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io"
	"io/ioutil"
	"net/http"
)

func callGetTicketDynamicField(ticketid string, username string, password string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetTicketDynamicField").GetString("uri")
	action := ConfObj.Sub("components.otrs.apis.GetTicketDynamicField").GetString("Action")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&Action=" + action + "&TicketID=" + ticketid + "&SessionID=" + sessionIDString
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
	return bodyText
}

//Function to get the details about ticket.

func (h *Handler) GetTicketDynamicField(w http.ResponseWriter, r *http.Request) {
	//body, _ := ioutil.ReadAll(r.Body)
	mapHttp := r.URL.Query()
	var ticketid string
	var username string
	var password string
	for key, value := range mapHttp {
		if key == "TicketID" {
			for _, valueStrg := range value {
				ticketid = valueStrg
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
	utils.ResponseAbstract(callGetTicketDynamicField(ticketid, username, password),w)

}
