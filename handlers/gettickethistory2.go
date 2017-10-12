package handlers

import (
	"github.com/Unotechsoftware/Hydrav3/utils"
	"github.com/Unotechsoftware/Hydrav3/lerna"
	"net/http"
)

func callGetTicketHistoryVersionTwo(username string, password string, ticketid string, pagesize string, pagenumber string) []uint8{
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.gettickethistoryversiontwo").GetString("uri")
	felicityaction := ConfObj.Sub("components.otrs.apis.gettickethistoryversiontwo.parameters").GetString("Action")
	url := felicitybaseurl + felicityapiuri + "?TicketID=" + ticketid + "&UserLogin=" + username + "&Password=" + password + "&Action=" + felicityaction + "&PageSize=" + pagesize + "&PageNumber=" + pagenumber
	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that creates a GET API that returns the history associated with a ticket.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ identifies the Ticket and __PageSize__ and __PageNumber__ that identifies the Page Limits to return a paginated list of history items linked to a ticket.
//
// Returns data as shown in examples.
func (h *Handler) GetTicketHistoryVersionTwo(w http.ResponseWriter, r *http.Request) {

	mapHttp := r.URL.Query()
	var userName string
	var password string
	var ticketid string
	var PageSize string
	var PageNumber string
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
		if key == "TicketID" {
			for _, valueStrg := range value {
				ticketid = valueStrg
			}
		}
		if key == "PageSize" {
			for _, valueStrg := range value {
				PageSize = valueStrg
			}
		}
		if key == "PageNumber" {
			for _, valueStrg := range value {
				PageNumber = valueStrg
			}
		}
	}

	utils.ResponseAbstract(callGetTicketHistoryVersionTwo(userName, password, ticketid, PageSize, PageNumber),w)

}
