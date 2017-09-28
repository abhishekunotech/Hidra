package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callWorkOrders(username string, password string, ticketid string) []uint8{

	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlistofworkorders").GetString("uri")
	ticketid = ConfObj.Sub("components.otrs.apis.getlistofworkorders.parameters").GetString("TicketId")
	url := felicitybaseurl + felicityapiuri + "?TicketID=" + ticketid + "&SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)
}

//Function to get list of work orders
// Request as http://ip-host/getListOfWorkOrders?ticketID=521&password=abhik&userLogin=abhik

func (h *Handler) GetListOfWorkOrders(w http.ResponseWriter, r *http.Request) {

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

	utils.ResponseAbstract(callWorkOrders(userName, password, ticketid),w)
}
