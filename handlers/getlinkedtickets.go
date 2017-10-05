package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callLinkedTickets(ticketid string, username string, password string) []uint8 {

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlinkedticketdetails").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?TicketID=" + ticketid + "&SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)

}

// This function is a handler that creates a GET API to get the Tickets Linked to a Ticket.
//
// **Business Logic**: Function takes as an input GET Parameter, __ticketID__ that will identify a ticket and return all the Tickets attached to it.
//
// Returns data as shown in examples.
func (h *Handler) GetLinkedTickets(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
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
	utils.ResponseAbstract(callLinkedTickets(ticketid, username, password),w)

}
