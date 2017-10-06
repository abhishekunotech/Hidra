package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)

// This Type defines the Input JSON Body for creating a POST request to link objects.
//
// It includes required parameters
type AgentTicketCompose_Request struct{
	UserLogin	string	`json:"UserLogin"`
	Password	string	`json:"Password"`
	Action	string	`json:"Action"`
	Subaction	string	`json:"Subaction"`
	TicketID	string	`json:"TicketID"`
	StateID		string	`json:"StateID"`
	From		string	`json:"From"`
	To		string	`json:"To"`
	Cc		string	`json:"Cc"`
	Bcc		string	`json:"Bcc"`
	Subject		string	`json:"Subject"`
	Body		string	`json:"Body"`
	InReplyTo	string	`json:"InReplyTo"`
	References	string	`json:"References"`
	ResponseID	string	`json:"ResponseID"`
	ReplyArticleID	string	`json:"ReplyArticleID"`
	ArticleTypeID	string	`json:"ArticleTypeID"`
	TimeUnits	string	`json:"TimeUnits"`
	Year		string	`json:"Year"`
	Month		string	`json:"Month"`
	Day		string	`json:"Day"`
	Hour		string	`json:"Hour"`
	Minute		string	`json:"Minute"`
	ReplyAll	string	`json:"ReplyAll"`
	CustomerTicketCounterToCustomer	string	`json:"CustomerTicketCounterToCustomer"`
	CustomerTicketText      	[]string        `json:"CustomerTicketText"`
	CustomerKey     		[]string        `json:"CustomerKey"`
	CustomerQueue  	 		[]string        `json:"CustomerQueue"`
	CustomerTicketCounterCcCustomer string  `json:"CustomerTicketCounterCcCustomer"`
	CcCustomerTicketText      	[]string        `json:"CcCustomerTicketText"`
	CcCustomerKey     		[]string        `json:"CcCustomerKey"`
	CustomerQueueCc  		[]string        `json:"CustomerQueueCc"`
	CustomerTicketCounterBccCustomer string  `json:"CustomerTicketCounterBccCustomer"`
	BccCustomerTicketText      	[]string        `json:"BccCustomerTicketText"`
	BccCustomerKey			[]string	`json:"BccCustomerKey"`
	BccCustomerQueue		[]string	`json:"BccCustomerQueue"`
	Attachment        		[]ATCR_Attachment        `json:"Attachment"`

}



//
//
//
type ATCR_Attachment struct{
	Content	string	`json:"Content"`
	ContentType	string	`json:"ContentType"`
	Filename	string	`json:"Filename"`
}


// This function is a handler that creates a POST API for reply to article.
//
// **Business Logic**: Function takes as input a JSON Body.
//
// Returns result message as success or failed.
func (h *Handler) AgentTicketCompose(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t AgentTicketCompose_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()
	utils.ResponseAbstract(AgentTicketCompose(t),w)
}

func AgentTicketCompose(T AgentTicketCompose_Request) []uint8 {

	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.AgentTicketCompose").GetString("uri")

	url := felicitybaseurl + felicityapiuri 
	j, err := json.Marshal(T)

	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)
	
	return utils.MakeHTTPPostCall(url,b)

}
