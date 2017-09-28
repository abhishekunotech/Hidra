package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
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

//Function to get the details about ticket.
// Request as http://ip-host/getTicketDetails?ticketID=521&password=abhik&userLogin=abhik
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
	utils.ResponseAbstract(callTicketDetails(userName, password, ticketid),w)
}
