package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callTicketAll(username string, password string, ticketid string) []uint8{

	sessionIDString := callCustomerSessionDetails(username, password)

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.TicketAll").GetString("uri")
	filter := ConfObj.Sub("components.otrs.apis.TicketAll.parameters").GetString("filter")

	url := felicitybaseurl + felicityapiuri + "?Filter=" + filter + "&SessionID=" + sessionIDString

	return utils.MakeHTTPGetCall(url)

}

//Function to get list of work orders
// Request as http://ip-host/ticketAll?ticketID=521&password=abhik&userLogin=abhik

func (h *Handler) TicketAll(w http.ResponseWriter, r *http.Request) {

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
	}

	utils.ResponseAbstract(callTicketAll(userName, password, ticketid),w)

}
