package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)

// This function is a handler that creates a POST request for Bulk Action on Ticket
//
// **Business Logic**: Function takes as an input a JSON Body in Request to generate the response.
//
// Returns the resulting message as success or failed based on the response
func (h *Handler) SetAgentTicketBulk(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t  BulkTicket_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	utils.ResponseAbstract(setAgentTicketBulk(t),w)
}

func setAgentTicketBulk(T BulkTicket_Request) []uint8 {

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.SetAgentTicketBulk").GetString("uri")
	//sessionIDString := callSessionDetails(T.UserLogin,T.Password)

	//url := felicitybaseurl+felicityapiuri+"?SessionID="+sessionIDString
	url := felicitybaseurl + felicityapiuri
	j, err := json.Marshal(T)

	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)

	return utils.MakeHTTPPostCall(url,b)

}

   
