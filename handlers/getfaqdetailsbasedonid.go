package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callGetPublicFAQ(username string, password string, itemid string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetPublicFAQ").GetString("uri")
	sessionIDString := callSessionDetails(username, password)

	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&ItemID=" + itemid + "&SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that creates a GET API to get the details of Public FAQ from its ID.
//
// **Business Logic**: Function takes as an input GET Parameter, __ItemID__ that will identify a public FAQ and get details of the FAQ.
//
// Returns data as shown in examples
func (h *Handler) GetPublicFAQ(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var itemid string
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
		if key == "ItemID" {
                        for _, valueStrg := range value {
                                itemid = valueStrg
                        }
                }

	}
	utils.ResponseAbstract(callGetPublicFAQ(username, password, itemid),w)

}
