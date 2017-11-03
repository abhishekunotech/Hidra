package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)


type CreateTicketCustomer_Request struct{
	UserLogin	string	`json:"UserLogin,omitempty"`
	Password	string	`json:"Password,omitempty"`
	CustomerUserLogin	string	`json:"CustomerUserLogin,omitempty"`
	SessionID	string	`json:"SessionID,omitempty"`
	Ticket  	*CreateCustomerTicketBody       `json:"Ticket,omitempty"`
        Article 	*CreateCustomerTicketArticleBody      `json:"Article,omitempty"`
	Attachment	*[]CTCR_Attachment	`json:"Attachment,omitempty"`
	DynamicField	*[]CTCR_DynamicField	`json:"DynamicField,omitempty"`
}


type CreateCustomerTicketBody struct{
        Title   string  `json:"Title,omitempty"`
        QueueID int     `json:"QueueID,omitempty"`
        Queue   string  `json:"Queue,omitempty"`
        LockID  int     `json:"LockID,omitempty"`
        Lock    string  `json:"Lock,omitempty"`
        TypeID  int     `json:"TypeID,omitempty"`
        Type    string  `json:"Type,omitempty"`
        ServiceID       int     `json:"ServiceID,omitempty"`
        Service string  `json:"Service,omitempty"`
        SLAID   int     `json:"SLAID,omitempty"`
        SLA     string  `json:"SLA,omitempty"`
        StateID int     `json:"StateID,omitempty"`
        State   string  `json:"State,omitempty"`
        PriorityID      int     `json:"PriorityID,omitempty"`
        Priority        string  `json:"Priority,omitempty"`
        OwnerID int     `json:"OwnerID,omitempty"`
        Owner   string  `json:"Owner,omitempty"`
        ResponsibleID   int     `json:"ResponsibleID,omitempty"`
        Responsible     string  `json:"Responsible,omitempty"`
        CustomerUser    string  `json:"CustomerUser,omitempty"`
        PendingTime     *PendingTimeBody `json:"PendingTime,omitempty"`
}

type CreateCustomerTicketArticleBody struct{
        ArticleTypeID   int     `json:"ArticleTypeID,omitempty"`
        ArticleType     string  `json:"ArticleType,omitempty"`
        SenderTypeID    int     `json:"SenderTypeID,omitempty"`
        SenderType      string  `json:"SenderType,omitempty"`
        AutoResponseType        string  `json:"AutoResponseType,omitempty"`
        From    string  `json:"From,omitempty"`
        Subject string  `json:"Subject,omitempty"`
        Body    string  `json:"Body,omitempty"`
        ContentType     string  `json:"ContentType,omitempty"`
        MimeType        string  `json:"MimeType,omitempty"`
        Charset string  `json:"Charset,omitempty"`
        HistoryType     string  `json:"HistoryType,omitempty"`
        HistoryComment  string  `json:"HistoryComment,omitempty"`
        TimeUnit        int     `json:"TimeUnit,omitempty"`
        NoAgentNotify   int     `json:"NoAgentNotify,omitempty"`
        ForceNotificationToUserID       []int   `json:"ForceNotificationToUserID,omitempty"`
        ExcludeNotificationToUserID     []int   `json:"ExcludeNotificationToUserID,omitempty"`
}

type CTCR_Attachment struct{
	Content	string	`json:"Content,omitempty"`
	ContentType	string	`json:"ContentType,omitempty"`
	Filename	string	`json:"Filename,omitempty"`

}

type CTCR_DynamicField	struct{
	Name	string	`json:"Name,omitempty"`
	Value	string	`json:"Value,omitempty"`
}

// This function is a handler that creates a POST request to update Ticket Common Agent Functions
//
// **Business Logic**: Function takes as an input a JSON Body and generates the response
//
// Returns data as shown in examples
func (h *Handler) PostCreateTicketCustomer(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t CreateTicketCustomer_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	utils.ResponseAbstract(postCreateTicketCustomer(t),w)
}

func postCreateTicketCustomer(T CreateTicketCustomer_Request) []uint8{

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.postcreateTicketcustomer").GetString("uri")

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
