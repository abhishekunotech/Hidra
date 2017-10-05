package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callGetGlobalSearch(username string, password string, term string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetGlobalSearch").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&term=" + term + "&SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that creates a GET API to search for a term in the Service Catalog
//
// **Business Logic**: Function takes as an input GET Parameter, __term__ that will search for the value of that parameter within the Admin Catalog.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetGlobalSearch(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var username string
	var password string
	var term string
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
		if key == "term" {
                        for _, valueStrg := range value {
                                term = valueStrg
                        }
                }

	}
	utils.ResponseAbstract(callGetGlobalSearch(username, password, term),w)

}
