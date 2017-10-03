package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callListOfLinkedFAQS(username string, password string, ticketid string) []uint8{
	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetListOfFAQs").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?TicketID=" + ticketid + "&SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)
}



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
