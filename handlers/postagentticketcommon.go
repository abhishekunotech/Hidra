package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)


type AgentTicketCommon_Request struct{
	UserLogin	string	`json:"UserLogin,omitempty"`
	Password	string	`json:"Password,omitempty"`
	TicketIDs	string	`json:"TicketIDs,omitempty"`
	Action	string	`json:"Action,omitempty"`
	Subaction	string	`json:"Subaction,omitempty"`
	TimeUnits	string	`json:"TimeUnits,omitempty"`
	NewStateID	string	`json:"NewStateID,omitempty"`
	CreateArticle	string	`json:"CreateArticle,omitempty"`
	ArticleTypeID	string	`json:"ArticleTypeID,omitempty"`
	Body	string	`json:"Body,omitempty"`
	Subject	string	`json:"Subject,omitempty"`
	DynamicField_ITSMReviewRequired	string	`json:"DynamicField_ITSMReviewRequired,omitempty"`
	NewOwnerID	string	`json:"NewOwnerID,omitempty"`
	PriorityRC	string	`json:"PriorityRC,omitempty"`
	ElementChanged	string	`json:"ElementChanged,omitempty"`
	Hour	string	`json:"Hour,omitempty"`
	Year	string	`json:"Year,omitempty"`
	Day	string	`json:"Day,omitempty"`
	Minute	string	`json:"Minute,omitempty"`
	Month	string	`json:"Month,omitempty"`
	ReplyToArticle  string  `json:"ReplyToArticle,omitempty"`
    UserListWithoutSelection    string  `json:"UserListWithoutSelection,omitempty"`
    NewPriorityID	string	`json:"NewPriorityID,omitempty"`
	TypeID	string	`json:"TypeID,omitempty"`
	NewResponsibleID	string	`json:"NewResponsibleID,omitempty"`
	Article	*[]ATCR_Article	`json:"Attachment,omitempty"`
	DynamicField	*[]ATCR_DynamicField	`json:"DynamicField,omitempty"`
}


type ATCR_Article struct{
	Content	string	`json:"Content,omitempty"`
	ContentType	string	`json:"ContentType,omitempty"`
	Filename	string	`json:"Filename,omitempty"`

}

type ATCR_DynamicField	struct{
	Name	string	`json:"Name,omitempty"`
	Value	string	`json:"Value,omitempty"`
}

// This function is a handler that creates a POST request to update Ticket Common Agent Functions
//
// **Business Logic**: Function takes as an input a JSON Body and generates the response
//
// Returns data as shown in examples
func (h *Handler) PostAgentTicketCommon(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t AgentTicketCommon_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	utils.ResponseAbstract(postAgentTicketCommon(t),w)
}

func postAgentTicketCommon(T AgentTicketCommon_Request) []uint8{

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.agentticketcommonpost").GetString("uri")

	url := felicitybaseurl + felicityapiuri 
	j, err := json.Marshal(T)
//	logger.Info(T)
	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)

	return utils.MakeHTTPPostCall(url,b)

}
