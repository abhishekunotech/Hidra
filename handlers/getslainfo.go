package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callSLAInfo(username string, password string, ticketid string) []uint8{

	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getslainfo").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString + "&TicketID=" + ticketid

	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that creates a GET API that returns the SLA information.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ identifies the Ticket and returns the information of the SLA assigned to the ticket.
//
// Returns data as shown in examples.
func (h *Handler) GetSLAInfo(w http.ResponseWriter, r *http.Request) {

	mapHttp := r.URL.Query()
	var userName string
	var password string
	var ticketid string
	for key, value := range mapHttp {
		if key == "userLogin" {
			for _, valueStrg := range value {
				userName = valueStrg
			}
		}
		if key == "password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
		if key == "ticketID" {
			for _, valueStrg := range value {
				ticketid = valueStrg
			}
		}
	}

	utils.ResponseAbstract(callSLAInfo(userName, password, ticketid),w)

}
