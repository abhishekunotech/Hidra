package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callGetCreatedTimeRange(username string, password string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetCreatedTimeRange").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri +  "?SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)
}


func (h *Handler) GetCreatedTimeRange(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var username string
	var password string
	for key, value := range mapHttp {
		if key == "userLogin" {
			for _, valueStrg := range value {
				username = valueStrg
			}
		}
		if key == "password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
	}
	utils.ResponseAbstract(callGetCreatedTimeRange(username, password),w)

}
