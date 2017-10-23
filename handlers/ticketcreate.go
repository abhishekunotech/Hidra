package handlers

import (
	"bytes"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"net/http"
	"time"
)
//The following struct provides the JSON structure to create a ticket.
type TicketResponseBody struct {
	ArticleID    string `json:"ArticleID,omitempty"`
	TicketNumber string `json:"TicketNumber"`
	TicketID     string `json:"TicketID"`
}
// This function is a handler that creates a POST API to create a ticket
//
// **Business Logic**: Function takes Article ID, Ticket Number and Ticket ID as POST parameters and creates a request to generate ticket.
//
// Returns the resulting message as success or failed based on the response.
func (h *Handler) Ticketcreate(w http.ResponseWriter, r *http.Request) {

	//ReadAll reads from response until an error or EOF and returns the data it read.
	bodyVal, err := ioutil.ReadAll(r.Body)

	if err != nil {
		logger.Error("Error Occured with Reading Body")
		logger.Error(err.Error())
	}

	bodyValStrg := string(bodyVal)

	//Function call to create ticket and get the response
	creatorOfTickets(bodyValStrg, w, r)

}

func creatorOfTickets(jsonInput string, w http.ResponseWriter, r *http.Request) {

	http.DefaultClient.Timeout = 10 * time.Second
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.ticketcreate").GetString("uri")
	passwordStrg := ConfObj.Sub("components.otrs.apis.ticketcreate.parameters").GetString("Password")
	usernameStrg := ConfObj.Sub("components.otrs.apis.ticketcreate.parameters").GetString("Username")
	sessionIDString := callSessionDetails(usernameStrg, passwordStrg)

	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString

	jsonStr1 := bytes.NewBufferString(jsonInput)

	bodyText := utils.MakeHTTPPostCall(url,jsonStr1)

	utils.ResponseAbstract(bodyText,w)
}
