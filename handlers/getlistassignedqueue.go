package handlers

import (
	"github.com/Unotechsoftware/Hydrav2/utils"
	"github.com/Unotechsoftware/Hydrav2/lerna"
	"net/http"
)

func callAssignedQueue(username string, password string) []uint8 {
	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlistassignedqueue").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)

}

// This function is a handler
//
// **Business Logic**: To be done.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListAssignedQueue(w http.ResponseWriter, r *http.Request) {

	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
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
	}

	assignedqueues := callAssignedQueue(userName, password)
	utils.ResponseAbstract(assignedqueues, w)
}
