package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callGetProcessInformation(ticketid string, username string, password string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetProcessInformation").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?TicketID=" + ticketid + "&SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)
}

//Function to get the details about ticket.

func (h *Handler) GetProcessInformation(w http.ResponseWriter, r *http.Request) {
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
	utils.ResponseAbstract(callGetProcessInformation(ticketid, username, password),w)

}
