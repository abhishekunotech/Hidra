package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callAgents(username string, password string, search string, term string) []uint8{

	sessionIDString := callSessionDetails(username, password)

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlistagents").GetString("uri")

	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString + "&Search=" + search + "&term=" + term
bodyText := utils.MakeHTTPGetCall(url)
	return bodyText
}

// This function is a handler that creates a GET API to search for an agent in the system
//
// **Business Logic**: Function takes as an input GET Parameter, __term__ that will search for agents whose login names match the parameters.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListAgents(w http.ResponseWriter, r *http.Request) {

	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	var search string
	var term string
	//var ticketid string
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
		if key == "Search" {
			for _, valueStrg := range value {
				search = valueStrg
			}
		}
		if key == "Term" {
			for _, valueStrg := range value {
				term = valueStrg
			}
		}
	}

	utils.ResponseAbstract(callAgents(userName, password, search, term),w)

}
