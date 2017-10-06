package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callQueueTemplateList(username string, password string, queueid string) []uint8 {

	sessionIDString := callSessionDetails(username, password)

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.queuetemplatelist").GetString("uri")

	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString + "&QueueID=" + queueid

	return utils.MakeHTTPGetCall(url)
}


func (h *Handler) GetQueueTemplateList(w http.ResponseWriter, r *http.Request) {

	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	var queueid string
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
		if key == "QueueID" {
			for _, valueStrg := range value {
				queueid = valueStrg
			}
		}
	}

	utils.ResponseAbstract(callQueueTemplateList(userName, password, queueid),w)

}
