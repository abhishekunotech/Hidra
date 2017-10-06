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
	UserLogin	string	`json:"UserLogin"`
	Password	string	`json:"Password"`
	SourceObject	string	`json:"SourceObject"`
	SourceKey	string	`json:"SourceKey"`
	TargetIdentifier       string  `json:"TargetIdentifier"`
	TypeIDs		[]string	`json:"TypeIDs"`
	StateIDs 	[]string        `json:"StateIDs"`
	TicketTitle	string	`json:"TicketTitle"`
	TicketNumber	string	`json:"TicketNumber"`
	PriorityIDs     []string        `json:"PriorityIDs"`
	TicketFulltext    string  `json:"TicketFulltext"`
	WorkOrderTitle    string  `json:"WorkOrderTitle"`
	ChangeTitle     string  `json:"ChangeTitle"`
	ChangeNumber    string  `json:"ChangeNumber"`
	Number          string  `json:"Number"`
	What    	string  `json:"What"`
	Title    	string  `json:"Title"`
	WorkOrderNumber    string  `json:"WorkOrderNumber"`
	InciStateIDs    []string        `json:"InciStateIDs"`
	DeplStateIDs    []string        `json:"DeplStateIDs"`
	Name    	string  `json:"Name"`
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
	felicityapiuri := ConfObj.Sub("components.otrs.apis.LinkObjectSearch").GetString("uri")

	url := felicitybaseurl + felicityapiuri 
	j, err := json.Marshal(T)

	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)
	return utils.MakeHTTPPostCall(url,b)

}
