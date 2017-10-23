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
	UserLogin	string	`json:"UserLogin,omitempty"`
	Password	string	`json:"Password,omitempty"`
	Action	string	`json:"Action,omitempty"`
	Subaction	string	`json:"Subaction,omitempty"`
	TicketID	string	`json:"TicketID,omitempty"`
	StateID		string	`json:"StateID,omitempty"`
	From		string	`json:"From,omitempty"`
	To		string	`json:"To,omitempty"`
	Cc		string	`json:"Cc,omitempty"`
	Bcc		string	`json:"Bcc,omitempty"`
	Subject		string	`json:"Subject,omitempty"`
	Body		string	`json:"Body,omitempty"`
	InReplyTo	string	`json:"InReplyTo,omitempty"`
	References	string	`json:"References,omitempty"`
	ResponseID	string	`json:"ResponseID,omitempty"`
	ReplyArticleID	string	`json:"ReplyArticleID,omitempty"`
	ArticleTypeID	string	`json:"ArticleTypeID,omitempty"`
	TimeUnits	string	`json:"TimeUnits,omitempty"`
	Year		string	`json:"Year,omitempty"`
	Month		string	`json:"Month,omitempty"`
	Day		string	`json:"Day,omitempty"`
	Hour		string	`json:"Hour,omitempty"`
	Minute		string	`json:"Minute,omitempty"`
	ReplyAll	string	`json:"ReplyAll,omitempty"`
	CustomerTicketCounterToCustomer	string	`json:"CustomerTicketCounterToCustomer,omitempty"`
	CustomerTicketText      	[]string        `json:"CustomerTicketText,omitempty"`
	CustomerKey     		[]string        `json:"CustomerKey,omitempty"`
	CustomerQueue  	 		[]string        `json:"CustomerQueue,omitempty"`
	CustomerTicketCounterCcCustomer string  `json:"CustomerTicketCounterCcCustomer,omitempty"`
	CcCustomerTicketText      	[]string        `json:"CcCustomerTicketText,omitempty"`
	CcCustomerKey     		[]string        `json:"CcCustomerKey,omitempty"`
	CustomerQueueCc  		[]string        `json:"CustomerQueueCc,omitempty"`
	CustomerTicketCounterBccCustomer string  `json:"CustomerTicketCounterBccCustomer,omitempty"`
	BccCustomerTicketText      	[]string        `json:"BccCustomerTicketText,omitempty"`
	BccCustomerKey			[]string	`json:"BccCustomerKey,omitempty"`
	BccCustomerQueue		[]string	`json:"BccCustomerQueue,omitempty"`
	Attachment        		*[]ATCR_Attachment        `json:"Attachment,omitempty"`

}



//
//
//
type ATCR_Attachment struct{
	Content	string	`json:"Content,omitempty"`
	ContentType	string	`json:"ContentType,omitempty"`
	Filename	string	`json:"Filename,omitempty"`
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
