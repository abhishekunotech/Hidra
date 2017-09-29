package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"net/http"
)

type PendingTimeBody struct{
	Year	int	`json:"Year,omitempty"`
	Month	int	`json:"Month,omitempty"`
	Day	int	`json:"Day,omitempty"`
	Hour	int	`json:"Hour,omitempty"`
	Minute	int	`json:"Minute,omitempty"`
	Diff	int	`json:"Diff,omitempty"`
}

type UpdateTicketBody struct{
	Title	string	`json:"Title,omitempty"`
	QueueID	int	`json:"QueueID,omitempty"`
	Queue	string	`json:"Queue,omitempty"`
	LockID	int	`json:"LockID,omitempty"`
	Lock	string	`json:"Lock,omitempty"`
	TypeID	int	`json:"TypeID,omitempty"`
	Type	string	`json:"Type,omitempty"`
	ServiceID	int	`json:"ServiceID,omitempty"`
	Service	string	`json:"Service,omitempty"`
	SLAID	int	`json:"SLAID,omitempty"`
	SLA	string	`json:"SLA,omitempty"`
	StateID	int	`json:"StateID,omitempty"`
	State	string	`json:"State,omitempty"`
	PriorityID	int	`json:"PriorityID,omitempty"`
	Priority	string	`json:"Priority,omitempty"`
	OwnerID	int	`json:"OwnerID,omitempty"`
	Owner	string	`json:"Owner,omitempty"`
	ResponsibleID	int	`json:"ResponsibleID,omitempty"`
	Responsible	string	`json:"Responsible,omitempty"`
	CustomerUser	string	`json:"CustomerUser,omitempty"`
	PendingTime	PendingTimeBody `json:"PendingTime,omitempty"`
}

type UpdateArticleBody struct{
	ArticleTypeID	int	`json:"ArticleTypeID,omitempty"`
	ArticleType	string	`json:"ArticleType,omitempty"`
	SenderTypeID	int	`json:"SenderTypeID,omitempty"`
	SenderType	string	`json:"SenderType,omitempty"`
	AutoResponseType	string	`json:"AutoResponseType,omitempty"`
	From	string	`json:"From,omitempty"`
	Subject	string	`json:"Subject,omitempty"`
	Body	string	`json:"Body,omitempty"`
	ContentType	string	`json:"ContentType,omitempty"`
	MimeType	string	`json:"MimeType,omitempty"`
	Charset	string	`json:"Charset,omitempty"`
	HistoryType	string	`json:"HistoryType,omitempty"`
	HistoryComment	string	`json:"HistoryComment,omitempty"`
	TimeUnit	int	`json:"TimeUnit,omitempty"`
	NoAgentNotify	int	`json:"NoAgentNotify,omitempty"`
	ForceNotificationToUserID	[]int	`json:"ForceNotificationToUserID,omitempty"`
	ExcludeNotificationToUserID	[]int	`json:"ExcludeNotificationToUserID,omitempty"`
}

type UpdateDynamicField struct{
	Name	string	`json:"Name,omitempty"`
	Value	string	`json:"Value,omitempty"`
}

type UpdateAttachment struct{
	Content	string	`json:"Content,omitempty"`
	ContentType	string	`json:"ContentType,omitempty"`
	FileName	string	`json:"Filename,omitempty"`
}

type UpdateTicket_Request struct {
	UserLogin      string   `json:"UserLogin,omitempty"`
	CustomerUserLogin	string	`json:"CustomerUserLogin,omitempty"`
	SessionID	int	`json:"SessionID,omitempty"`
	Password       string   `json:"Password,omitempty"`
	TicketID	int	`json:"TicketID,omitempty"`
	TicketNumber	string	`json:"TicketNumber,omitempty"`
	Ticket	UpdateTicketBody	`json:"Ticket,omitempty"`
	//Article	UpdateArticleBody	`json:"Article,omitempty"`
	DynamicField	[]UpdateDynamicField	`json:"DynamicField,omitempty"`
	Attachment	[]UpdateAttachment	`json:"Attachment,omitempty"`
}

func (h *Handler) UpdateTicketInfo(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t UpdateTicket_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	utils.ResponseAbstract(updateTicketinfo(t),w)
}

func updateTicketinfo(T UpdateTicket_Request) []uint8{

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.updateticketinfo").GetString("uri")

	url := felicitybaseurl + felicityapiuri 
	j, err := json.Marshal(T)
//	logger.Info(T)
	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, b)

	if err != nil {
		logger.Error("\n\n Request to Create Request Failed \n\n")
		logger.Error(err.Error())
	}

	logger.Info("Request")

	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("\n\n POST REQUEST TO FELICITY FAILED \n\n")
		logger.Error(err.Error())
	}
	//req.Close = true
	bodyText, err := ioutil.ReadAll(resp.Body)
	return bodyText

}
