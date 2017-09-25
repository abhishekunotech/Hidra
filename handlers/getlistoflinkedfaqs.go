package handlers

import (
	"github.com/antigloss/go/logger"

	"github.com/Unotechsoftware/Hydra/utils"

	"github.com/Unotechsoftware/Hydra/lerna"

	"net/http"

	"io/ioutil"
)

func callListOfLinkedFAQS(username string, password string, ticketid string) []uint8{

	sessionIDString := callSessionDetails(username, password)

	logger.Info("session id is ::", sessionIDString)

	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")

	logger.Info("base url:- ", felicitybaseurl)

	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetListOfFAQs").GetString("uri")


	url := felicitybaseurl + felicityapiuri + "?TicketID=" + ticketid + "&SessionID=" + sessionIDString


	res, err := http.Get(url)

	if err != nil {

		logger.Error(err.Error())

	}

	bodyText, err := ioutil.ReadAll(res.Body)

	return bodyText

}

//Function to get list of work orders

// Request as http://ip-host/getListOfWorkOrders?ticketID=521&password=abhik&userLogin=abhik

func (h *Handler) GetListOfLinkedFAQs(w http.ResponseWriter, r *http.Request) {


	mapHttp := utils.RequestAbstractGet(r)

	var userName string

	var password string

	var ticketid string

	for key, value := range mapHttp {

		if key == "TicketID" {

			for _, valueStrg := range value {

				ticketid = valueStrg

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

	LinkedFaqsList := callListOfLinkedFAQS(userName, password, ticketid)
	utils.ResponseAbstract(LinkedFaqsList, w)

}
