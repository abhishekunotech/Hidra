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
