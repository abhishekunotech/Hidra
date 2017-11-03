package handlers

import (
	"fmt"
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)

// This Type defines the Input JSON Body for creating a request to close the ticket
//
// It includes required parameters
type Close_Request struct{
	UserLogin	string	`json:"UserLogin"`
	Password	string	`json:"Password"`
	Action		string	`json:"Action"`
	TicketIDs	[]string	`json:"TicketIDs"`
	Subaction	string	`json:"Subaction"`
	ArticleType	string	`json:"ArticleType"`
	Unlock		string	`json:"Unlock"`
	StateID		string	`json:"StateID"`
	Subject		string	`json:"Subject"`
	Body		string	`json:"Body"`
}

// This function is a handler that creates a request to close the ticket
//
// **Business Logic**: Function takes as an input a JSON Body and uses Ticket details in Request Body to generate a request
//
// Returns data as shown in examples
func (h *Handler) PostCloseAgentTicket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IN close tik")
	decoder := json.NewDecoder(r.Body)
	var t Close_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()
	utils.ResponseAbstract(PostCloseAgentTicket(t),w)
}

func PostCloseAgentTicket(T Close_Request) []uint8 {

	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.postcloseagentticket").GetString("uri")

	url := felicitybaseurl + felicityapiuri 
	j, err := json.Marshal(T)

	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)
	
	return utils.MakeHTTPPostCall(url,b)

}
