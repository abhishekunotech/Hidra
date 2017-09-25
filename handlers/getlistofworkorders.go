package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"net/http"
)

func callWorkOrders(username string, password string, ticketid string) []uint8{

	sessionIDString := callSessionDetails(username, password)

	logger.Info("session id is ::", sessionIDString)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	logger.Info("base url:- ", felicitybaseurl)
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlistofworkorders").GetString("uri")
	ticketid = ConfObj.Sub("components.otrs.apis.getlistofworkorders.parameters").GetString("TicketId")

	url := felicitybaseurl + felicityapiuri + "?TicketID=" + ticketid + "&SessionID=" + sessionIDString

	logger.Info("url is::",url)
	res, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
	}

	bodyText, err := ioutil.ReadAll(res.Body)
	return bodyText
}

//Function to get list of work orders
// Request as http://ip-host/getListOfWorkOrders?ticketID=521&password=abhik&userLogin=abhik

func (h *Handler) GetListOfWorkOrders(w http.ResponseWriter, r *http.Request) {

	mapHttp := r.URL.Query()
	var userName string
	var password string
	var ticketid string
	for key, value := range mapHttp {
		if key == "ticketID" {
			for _, valueStrg := range value {
				ticketid = valueStrg
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

	utils.ResponseAbstract(callWorkOrders(userName, password, ticketid),w)
}
