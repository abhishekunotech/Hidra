package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)

type Pickup_Request struct{
	UserLogin	string	`json:"UserLogin"`
	Password	string	`json:"Password"`
	Action		string	`json:"Action"`
	TicketIDs	[]string	`json:"TicketIDs"`
	Subaction	string	`json:"Subaction"`
	ArticleType	string	`json:"ArticleType"`
	Unlock		string	`json:"Unlock"`
	OwnerID		string	`json:"OwnerID"`
	Subject		string	`json:"Subject"`
	Body		string	`json:"Body"`
}

func (h *Handler) PostPickupAgentTicket(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t Pickup_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	utils.ResponseAbstract(pickupAgentTicket(t),w)
}

func pickupAgentTicket(T Pickup_Request) []uint8 {

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.postpickupagentticket").GetString("uri")
	url := felicitybaseurl + felicityapiuri  
	j, err := json.Marshal(T)

	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)

	return utils.MakeHTTPPostCall(url,b)
}