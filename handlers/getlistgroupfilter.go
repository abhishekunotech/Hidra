package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callGroupFilter(username string, password string, UserAccess string) []uint8{
	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlistgroupfilter").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString + "&UserAccess=" + UserAccess
	return utils.MakeHTTPGetCall(url)
}


// This function is a handler that creates a GET API <TBD>
//
// **Business Logic**: Function takes as an input GET Parameter, __UserAccess__ <TBD>
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListGroupFilter(w http.ResponseWriter, r *http.Request) {

	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	var userAccess string
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
		if key == "UserAccess" {
			for _, valueStrg := range value {
				userAccess = valueStrg
			}
		}
	}

	groupfilter := callGroupFilter(userName, password, userAccess)
	utils.ResponseAbstract(groupfilter, w)
}
