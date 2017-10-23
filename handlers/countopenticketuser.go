package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callCountOfOpenTicketsCustomerUser(custID string, username string, password string, custuser string) []uint8{

	sessionIDString := callSessionDetails(username, password)

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.CountOfOpenTicketsUser").GetString("uri")
	state := ConfObj.Sub("components.otrs.apis.CountOfOpenTicketsUser.parameters").GetString("state")
	url := felicitybaseurl + felicityapiuri + "?State=" + state + "&SessionID=" + sessionIDString + "&CustomerID=" + custID + "&CustomerUser=" + custuser

	return utils.MakeHTTPGetCall(url)
}


// This function is a handler that decodes the request from UI
//
// **Business Logic**: Function takes as input a JSON Body and assigns the Tickets defined by TicketIDs to the OwnerID defined in the JSON Request
//
// Returns the ticketID and ticketNumber
func (h *Handler) GetCountOfOpenTicketsCustomerUser(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)

	var userName string
	var password string
	var custID string
	var custUser string
	for key, value := range mapHttp {
		if key == "CustomerID" {
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
		if key == "CustomerUser" {
			for _, valueStrg := range value {
				custUser = valueStrg
			}

		}
	}

	utils.ResponseAbstract(callCountOfOpenTicketsCustomerUser(custID, userName, password, custUser),w)

}
