package handlers

import (
	"github.com/Unotechsoftware/Hydrav3/utils"
	"github.com/Unotechsoftware/Hydrav3/lerna"
	"net/http"
)

func callListOfLinkedChange(username string, password string, ticketid string) []uint8{

	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlistoflinkedchange").GetString("uri")
	ticketid = ConfObj.Sub("components.otrs.apis.getlistoflinkedchange.parameters").GetString("TicketId")

	url := felicitybaseurl + felicityapiuri + "?TicketID=" + ticketid + "&SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)

}

// This function is a handler that creates a GET API that returns the list of Changes linked to a ticket
//
// **Business Logic**: Function takes as an input GET Parameter, __ticketID__ that identifies a ticket and returns Linked Changes
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListOfLinkedChange(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	var ticketid string
	for key, value := range mapHttp {
		if key == "ticketID" {
			for _, valueStrg := range value {
				ticketid = valueStrg
			}
		}
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

	linkedChangeList := callListOfLinkedChange(userName, password, ticketid)
	utils.ResponseAbstract(linkedChangeList, w)
}
