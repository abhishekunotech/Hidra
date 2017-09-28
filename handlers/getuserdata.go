package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func getuserdata(username string, password string, userid string) []uint8{

	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetUserData").GetString("uri")

	url := felicitybaseurl + felicityapiuri + "?UserID=" + userid + "&SessionID=" + sessionIDString


	return utils.MakeHTTPGetCall(url)

}

//Function to get list of work orders
// Request as http://ip-host/getListOfWorkOrders?ticketID=521&password=abhik&userLogin=abhik

func (h *Handler) GetUserData(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	var userid string
	for key, value := range mapHttp {
		if key == "UserID" {
			for _, valueStrg := range value {
				userid = valueStrg
			}
		}
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
	}

	ciList := getuserdata(userName, password, userid)
	utils.ResponseAbstract(ciList,w)
}
