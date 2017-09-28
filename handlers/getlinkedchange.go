package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)


func callLinkedChanges(username string, password string, ticketid string) []uint8{
	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlinkedchanges").GetString("uri")
	felicityticketid := ConfObj.Sub("components.otrs.apis.getlinkedchanges.parameters").GetString("TicketID")
	url := felicitybaseurl + felicityapiuri + "?TicketID=" + felicityticketid + "&SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)
}

//Function to get the details about ticket.
// Request as http://ip-host/getLinkedChange?ticketID=627&password=abhik&username=abhik
func (h *Handler) GetLinkedChange(w http.ResponseWriter, r *http.Request) {
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
		if key == "username" {
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
	utils.ResponseAbstract(callLinkedChanges(userName, password, ticketid),w)
}
