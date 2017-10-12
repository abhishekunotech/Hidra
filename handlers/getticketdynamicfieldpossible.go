
package handlers

import (
	"github.com/Unotechsoftware/Hydrav3/utils"
	"github.com/Unotechsoftware/Hydrav3/lerna"
	"net/http"
)

func callGetTicketDynamicFieldPossible(ticketid string, username string, password string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetTicketDynamicFieldPossible").GetString("uri")
	action := ConfObj.Sub("components.otrs.apis.GetTicketDynamicFieldPossible").GetString("Action")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&Action=" + action + "&TicketID=" + ticketid + "&SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that creates a GET API that returns a list of dynamic fields attached to a Ticket and their corresponding values.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ identifies the Ticket and returns a list of Dynamic Field names and corresponding values.
//
// Returns data as shown in examples.
func (h *Handler) GetTicketDynamicFieldPossible(w http.ResponseWriter, r *http.Request) {
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
