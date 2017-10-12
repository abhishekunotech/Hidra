package handlers

import (
	"github.com/Unotechsoftware/Hydrav2/utils"
	"github.com/Unotechsoftware/Hydrav2/lerna"
	"net/http"
)

func callGetLinkTypeList(username string, password string, sourceobj string, targetid string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetLinkTypeList").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&SourceObject=" + sourceobj + "&TargetIdentifier=" + targetid + "&SessionID=" + sessionIDString 
	
	return utils.MakeHTTPGetCall(url)

}

// This function is a handler that creates a GET API to get details about Link Type List. 
//
// **Business Logic**: Function takes as an input GET Parameters UserLogin, Password, SourceObject and TargetIdentifier and generate the response
//
// Returns data as shown in examples
func (h *Handler) GetLinkTypeList(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var username string
	var password string
	var sourceobj string
	var targetid string

	for key, value := range mapHttp {
		if key == "UserLogin" {
			for _, valueStrg := range value {
				username = valueStrg
			}
		}
		if key == "SourceObject" { 
                        for _, valueStrg := range value {
                                sourceobj = valueStrg
                        }
                }
		if key == "TargetIdentifier" { 
                        for _, valueStrg := range value {
                                targetid = valueStrg
                        }
                }

		if key == "Password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
	}
	utils.ResponseAbstract(callGetLinkTypeList(username, password, sourceobj, targetid),w)
}
