package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callListType(username string, password string) []uint8{

	sessionIDString := callSessionDetails(username, password)

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.listtype").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString

	return utils.MakeHTTPGetCall(url)

}

// This function is a handler that provides the type of associated ID.
//
// **Business Logic**: Function takes Username and Password in Request Body to generate response.
//
// Returns data as shown in examples
func (h *Handler) ListType(w http.ResponseWriter, r *http.Request) {

	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	for key, value := range mapHttp {
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

	utils.ResponseAbstract(callListType(userName, password),w)

}
