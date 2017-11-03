package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
	"fmt"
)

// This Type defines the Input JSON Body for creating a POST request.
//
// It includes required parameters
type LinkObjectSearch_Request struct{
	UserLogin	string	`json:"UserLogin,omitempty"`
	Password	string	`json:"Password,omitempty"`
	SourceObject	string	`json:"SourceObject,omitempty"`
	SourceKey	string	`json:"SourceKey,omitempty"`
	TargetIdentifier       string  `json:"TargetIdentifier,omitempty"`
	TypeIDs		[]string	`json:"TypeIDs,omitempty"`
	StateIDs 	[]string        `json:"StateIDs,omitempty"`
	TicketTitle	string	`json:"TicketTitle,omitempty"`
	TicketNumber	string	`json:"TicketNumber,omitempty"`
	PriorityIDs     []string        `json:"PriorityIDs,omitempty"`
	TicketFulltext    string  `json:"TicketFulltext,omitempty"`
	WorkOrderTitle    string  `json:"WorkOrderTitle,omitempty"`
	ChangeTitle     string  `json:"ChangeTitle,omitempty"`
	ChangeNumber    string  `json:"ChangeNumber,omitempty"`
	Number          string  `json:"Number,omitempty"`
	What    	string  `json:"What,omitempty"`
	Title    	string  `json:"Title,omitempty"`
	WorkOrderNumber    string  `json:"WorkOrderNumber,omitempty"`
	InciStateIDs    []string        `json:"InciStateIDs,omitempty"`
	DeplStateIDs    []string        `json:"DeplStateIDs,omitempty"`
	Name    	string  `json:"Name,omitempty"`
	TypeIdentifier	string	`json:"TypeIdentifier,omitempty"`
	LinkTargetKeys	[]string	`json:"LinkTargetKeys,omitempty"`
}

// This function is a handler that creates a POST API
//
// **Business Logic**: Function takes as input a JSON Body and links objects based on the parameters in the JSON Request.
//
// Returns result message
func (h *Handler) LinkObjectSearch(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t LinkObjectSearch_Request
	err := decoder.Decode(&t)
	fmt.Println("data",t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()
	utils.ResponseAbstract(LinkObjectSearch(t),w)
}

func LinkObjectSearch(T LinkObjectSearch_Request) []uint8 {

	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetLinkObjectSearch").GetString("uri")

	url := felicitybaseurl + felicityapiuri 
	j, err := json.Marshal(T)

	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)
	return utils.MakeHTTPPostCall(url,b)

}
