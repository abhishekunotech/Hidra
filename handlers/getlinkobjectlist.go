package handlers

import (
	"github.com/Unotechsoftware/Hydrav2/utils"
	"github.com/Unotechsoftware/Hydrav2/lerna"
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

// This function is a handler that creates a GET API to get linked object lists
//
// **Business Logic**: Function takes as an input GET Parameter UserLogin and password.
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
