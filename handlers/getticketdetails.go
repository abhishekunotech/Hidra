package handlers

import (
	"github.com/Unotechsoftware/Hydrav3/utils"
	"github.com/Unotechsoftware/Hydrav3/lerna"
	"net/http"
)

func callTicketDetails(username string, password string, ticketid string) []uint8{
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getticketdetails").GetString("uri")
	felicityusername := ConfObj.Sub("components.otrs.apis.getticketdetails.parameters").GetString("UserLogin")
	felicitypassword := ConfObj.Sub("components.otrs.apis.getticketdetails.parameters").GetString("Password")
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + felicityusername + "&Password=" + felicitypassword
	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that creates a GET API that returns the details of the Ticket.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ identifies the Ticket and returns the details of that ticket.
//
// Returns data as shown in examples.
func (h *Handler) GetTicketDetails(w http.ResponseWriter, r *http.Request) {
	mapHttp := r.URL.Query()
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
	utils.ResponseAbstract(callTicketDetails(userName, password, ticketid),w)
}
