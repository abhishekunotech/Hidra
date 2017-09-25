package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"net/http"
)

func callListType(username string, password string) []uint8{

	sessionIDString := callSessionDetails(username, password)

	logger.Info("session id is ::", sessionIDString)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	logger.Info("base url:- ", felicitybaseurl)
	felicityapiuri := ConfObj.Sub("components.otrs.apis.listtype").GetString("uri")
	logger.Info("API URI")
	logger.Info(felicityapiuri)
	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString

	res, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
	}

	bodyText, err := ioutil.ReadAll(res.Body)
	return bodyText

}

//Function to get list of work orders
// Request as http://ip-host/getListOfWorkOrders?ticketID=521&password=abhik&userLogin=abhik

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
