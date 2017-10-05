package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callUserColumnPreference(username string, password string, Action string) []uint8{

	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getusercolumnpreference").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?Action=" + Action + "&SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that provides the details about user column preferences based on action 
//
// **Business Logic**: Function uses Username and Password in Request Body to generate the response
//
// Returns data as shown in examples
func (h *Handler) GetUserColumnPreferences(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	var Action string
	for key, value := range mapHttp {
		if key == "Action" {
			for _, valueStrg := range value {
				Action = valueStrg
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

	utils.ResponseAbstract(callUserColumnPreference(userName, password, Action),w)

}
