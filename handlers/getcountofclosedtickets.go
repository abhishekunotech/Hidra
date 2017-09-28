package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callCountClosedTickets(username string, password string, customerid string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getcountclosedtickets").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?CustomerID=" + customerid + "&SessionID=" + sessionIDString + "&State=close"
	
	return utils.MakeHTTPGetCall(url)

}

//Function to get the details about ticket.

func (h *Handler) GetCountofClosedTickets(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var customerid string
	var username string
	var password string
	for key, value := range mapHttp {
		if key == "CustomerID" {
			for _, valueStrg := range value {
				customerid = valueStrg
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
	utils.ResponseAbstract(callCountClosedTickets(username, password, customerid),w)

}
