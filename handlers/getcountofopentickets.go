package handlers

import (
	"github.com/Unotechsoftware/Hydrav2/utils"
	"github.com/Unotechsoftware/Hydrav2/lerna"
	"net/http"
)

func callCountOfOpenTickets(custID string, username string, password string) []uint8 {

	sessionIDString := callSessionDetails(username, password)

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.countofopentickets").GetString("uri")
	state := ConfObj.Sub("components.otrs.apis.countofopentickets.parameters").GetString("state")

	url := felicitybaseurl + felicityapiuri + "?State=" + state + "&SessionID=" + sessionIDString + "&CustomerID=" + custID
	
	bodyText := utils.MakeHTTPGetCall(url)
	return bodyText
}

// This function is a handler that creates a GET API to get the count of open tickets that were created by/assigned to agent.
//
// **Business Logic**: Function that obtains a list of Tickets that are in "Open" state.
//
// Returns data as shown in examples
func (h *Handler) GetCountOfOpenTickets(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)

	var userName string
	var password string
	var custID string
	for key, value := range mapHttp {
		if key == "custID" {
			for _, valueStrg := range value {
				custID = valueStrg
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

	utils.ResponseAbstract(callCountOfOpenTickets(custID, userName, password),w)
}
