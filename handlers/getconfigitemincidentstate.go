package handlers

import (
	"github.com/Unotechsoftware/Hydrav2/utils"
	"github.com/Unotechsoftware/Hydrav2/lerna"
	"net/http"
)

func callGetConfigItemIncidentState(username string, password string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetConfigItemIncidentState").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&SessionID=" + sessionIDString 
	
	return utils.MakeHTTPGetCall(url)

}

// This function is a handler that creates a GET API to get Link Object CI Incident State Dropdown List. 
//
// **Business Logic**: Function takes as an input GET Parameter UserLogin and Password.
//
// Returns data as shown in examples
func (h *Handler) GetConfigItemIncidentState(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var username string
	var password string
	for key, value := range mapHttp {
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
	utils.ResponseAbstract(callGetConfigItemIncidentState(username, password),w)
}
