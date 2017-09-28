package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callUserColumnPreference(username string, password string, Action string) []uint8{

	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getusercolumnpreference").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?Action=" + Action + "&SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)
}

//Function to get list of work orders
// Request as http://ip-host/getListOfWorkOrders?ticketID=521&password=abhik&userLogin=abhik

func (h *Handler) GetUserColumnPreferences(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	var Action string
	for key, value := range mapHttp {
		if key == "Action" {
			for _, valueStrg := range value {
				Action = valueStrg
			}
		}
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

	utils.ResponseAbstract(callUserColumnPreference(userName, password, Action),w)

}
