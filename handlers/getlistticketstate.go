package handlers

import (
	"github.com/Unotechsoftware/Hydrav2/utils"
	"github.com/Unotechsoftware/Hydrav2/lerna"
	"net/http"
)

func callListTicketState(username string, password string) []uint8{

	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlistticketstate").GetString("uri")

	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that creates a GET API that returns a list of tickets and their corresponding states.
//
// **Business Logic**: Function returns a list of Tickets and their corresponding states.
//
// Returns data as shown in examples.
func (h *Handler) GetListTicketState(w http.ResponseWriter, r *http.Request) {

	mapHttp := r.URL.Query()
	var userName string
	var password string
	for key, value := range mapHttp {
		if key == "UserLogin" {
			for _, valueStrg := range value {
				userName = valueStrg
			}
		}
		if key == "Password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}

	}

	utils.ResponseAbstract(callListTicketState(userName, password),w)

}
