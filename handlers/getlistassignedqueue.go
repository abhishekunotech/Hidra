package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callAssignedQueue(username string, password string) []uint8 {
	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlistassignedqueue").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)

}


func (h *Handler) GetListAssignedQueue(w http.ResponseWriter, r *http.Request) {

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

	assignedqueues := callAssignedQueue(userName, password)
	utils.ResponseAbstract(assignedqueues, w)
}
