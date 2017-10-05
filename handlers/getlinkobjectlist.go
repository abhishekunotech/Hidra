package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callGetLinkObjectList(username string, password string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetLinkObjectList").GetString("uri")
	source := ConfObj.Sub("components.otrs.apis.GetLinkObjectList").GetString("SourceObject")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&SourceObject=" + source + "&SessionID=" + sessionIDString 
	
	return utils.MakeHTTPGetCall(url)

}

// This function is a handler that creates a GET API to get details about navigation bar of an agent on the basis of access.
//
// **Business Logic**: Function takes as an input GET Parameter UserLogin and Password that will identify the agent and obtain details of nav bar.
//
// Returns data as shown in examples
func (h *Handler) GetLinkObjectList(w http.ResponseWriter, r *http.Request) {
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
	utils.ResponseAbstract(callGetLinkObjectList(username, password),w)
}
