package handlers

import (
	"github.com/antigloss/go/logger"

	"encoding/json"

	"github.com/Unotechsoftware/Hydra/lerna"

	"net/http"

	"io/ioutil"
)

func callListOfLinkedFAQS(w http.ResponseWriter, r *http.Request, username string, password string, ticketid string) {

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

	var data interface{}

	err = json.Unmarshal(bodyText, &data)

	if err != nil {

		logger.Error(err.Error())

	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)


}

//Function to get list of work orders

// Request as http://ip-host/getListOfWorkOrders?ticketID=521&password=abhik&userLogin=abhik

func (h *Handler) GetListOfLinkedFAQs(w http.ResponseWriter, r *http.Request) {


	mapHttp := r.URL.Query()

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

	callListOfLinkedFAQS(w, r, userName, password, ticketid)


}
