package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"net/http"
)

type PendingTimeBody struct{
	Year	int	`json:"Year"`
	Month	int	`json:"Month"`
	Day	int	`json:"Day"`
	Hour	int	`json:"Hour"`
	Minute	int	`json:"Minute"`
	Diff	int	`json:"Diff"`
}

type UpdateTicketBody struct{
	Title	string	`json:"Title"`
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
	PriorityID	int	`json:"PriorityID"`
	Priority	string	`json:"Priority"`
	OwnerID	int	`json:"OwnerID,omitempty"`
	Owner	string	`json:"Owner,omitempty"`
	ResponsibleID	int	`json:"ResponsibleID,omitempty"`
	Responsible	string	`json:"Responsible,omitempty"`
	CustomerUser	string	`json:"CustomerUser"`
	PendingTime	PendingTimeBody `json:"PendingTime"`
}

type UpdateArticleBody struct{
	ArticleTypeID	int	`json:"ArticleTypeID,omitempty"`
	ArticleType	string	`json:"ArticleType,omitempty"`
	SenderTypeID	int	`json:"SenderTypeID,omitempty"`
	SenderType	string	`json:"SenderType,omitempty"`
	AutoResponseType	string	`json:"AutoResponseType"`
	From	string	`json:"From"`
	Subject	string	`json:"Subject"`
	Body	string	`json:"Body"`
	ContentType	string	`json:"ContentType"`
	MimeType	string	`json:"MimeType"`
	Charset	string	`json:"Charset"`
	HistoryType	string	`json:"HistoryType"`
	HistoryComment	string	`json:"HistoryComment"`
	TimeUnit	int	`json:"TimeUnit"`
	NoAgentNotify	int	`json:"NoAgentNotify"`
	ForceNotificationToUserID	[]int	`json:"ForceNotificationToUserID"`
	ExcludeNotificationToUserID	[]int	`json:"ExcludeNotificationToUserID"`
}

type UpdateDynamicField struct{
	Name	string	`json:"Name"`
	Value	string	`json:"Value"`
}

type UpdateAttachment struct{
	Content	string	`json:"Content"`
	ContentType	string	`json:"ContentType"`
	FileName	string	`json:"Filename"`
}

type UpdateTicket_Request struct {
	UserLogin      string   `json:"UserLogin"`
	CustomerUserLogin	string	`json:"CustomerUserLogin"`
	SessionID	int	`json:"SessionID"`
	Password       string   `json:"Password"`
	TicketID	int	`json:"TicketID"`
	TicketNumber	string	`json:"TicketNumber"`
	Ticket	UpdateTicketBody	`json:"Ticket"`
	Article	UpdateArticleBody	`json:"Article"`
	DynamicField	[]UpdateDynamicField	`json:"DynamicField"`
	Attachment	[]UpdateAttachment	`json:"Attachment"`
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

	updateTicketinfo(t, w, r)
}

func updateTicketinfo(T UpdateTicket_Request, w http.ResponseWriter, r *http.Request) {

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.updateticketinfo").GetString("uri")

	url := felicitybaseurl + felicityapiuri 
	j, err := json.Marshal(T)

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
	var data interface{}
	err = json.Unmarshal(bodyText, &data)
	if err != nil {
		logger.Error("\n\n Error Occured in unmarshalling Session JSON \n\n")
		logger.Error(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}
