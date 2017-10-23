package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)

// This Type defines the Input JSON Body for creating a POST request
//
// It includes required parameters
type AgentTicketForward_Request struct{
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
	ComposeStateID  string  `json:"ComposeStateID"`
	ArticleTypeID	string	`json:"ArticleTypeID"`
	TimeUnits	string	`json:"TimeUnits"`
	Year		string	`json:"Year"`
	Month		string	`json:"Month"`
	Day		string	`json:"Day"`
	Hour		string	`json:"Hour"`
	Minute		string	`json:"Minute"`
	ACLCompatGetParam      string  `json:"ACLCompatGetParam"`
	QueueID      string  `json:"QueueID"`
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
	Attachment        		[]ATFR_Attachment        `json:"Attachment"`

}



//
//
//
type ATFR_Attachment struct{
	Content	string	`json:"Content"`
	ContentType	string	`json:"ContentType"`
	Filename	string	`json:"Filename"`
}


// This function is a handler that creates a POST API for forward to article.
//
// **Business Logic**: Function takes as input a JSON Body.
//
// Returns result message as success or failed.
func (h *Handler) AgentTicketForward(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t AgentTicketForward_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()
	utils.ResponseAbstract(AgentTicketForward(t),w)
}

func AgentTicketForward(T AgentTicketForward_Request) []uint8 {

	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.AgentTicketForward").GetString("uri")

	url := felicitybaseurl + felicityapiuri 
	j, err := json.Marshal(T)

	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)
	
	return utils.MakeHTTPPostCall(url,b)

}
