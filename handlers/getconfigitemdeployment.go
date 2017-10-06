package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callGetConfigItemDeploymentState(username string, password string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetConfigItemDeploymentState").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&SessionID=" + sessionIDString 
	
	return utils.MakeHTTPGetCall(url)

}

// This function is a handler that creates a GET API to get Link Object CI Deployment State Dropdown List. 
//
// **Business Logic**: Function takes as an input GET Parameter UserLogin and Password.
//
// Returns data as shown in examples
func (h *Handler) GetConfigItemDeploymentState(w http.ResponseWriter, r *http.Request) {
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
	utils.ResponseAbstract(callGetConfigItemDeploymentState(username, password),w)
}
