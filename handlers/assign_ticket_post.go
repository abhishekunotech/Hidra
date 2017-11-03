/*
    This is the List of Handler Functions

        This package provides the source code for all the APIs exposed by Hydra.

        Author: Operations Management Team - Unotech Software
*/

package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)

// This structure defines a Request for assigning a ticket
type Assign_Request struct{
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

// This function is a handler that decodes the request from UI
//
// **Business Logic**: Function takes as input a JSON Body and assigns the Tickets defined by TicketIDs to the OwnerID defined in the JSON Request
//
// Returns the ticketID and ticketNumber
func (h *Handler) PostAssignAgentTicket(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t Assign_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	utils.ResponseAbstract(assignAgentTicket(t),w)
}


// This function defines the business logic and pulls data from configuration to call the API to assign a ticket to an agent
func assignAgentTicket(T Assign_Request) []uint8 {

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.postassignagentticket").GetString("uri")

	url := felicitybaseurl + felicityapiuri 
	j, err := json.Marshal(T)

	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)
	
	return utils.MakeHTTPPostCall(url,b)

}
