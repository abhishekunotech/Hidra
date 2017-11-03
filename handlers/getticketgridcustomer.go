package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydrav4/utils"
	"github.com/Unotechsoftware/Hydrav4/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)


type GetTicketGridCustomer_Request struct{
	CustomerUserLogin	string	`json:"CustomerUserLogin,omitempty"`
	Password	string	`json:"Password,omitempty"`
	StateID		[]int	`json:"StateIDs,omitempty"`
	PriorityIDs	[]int	`json:"PriorityIDs,omitempty"`
	PageSize	int	`json:"PageSize,omitempty"`
	PageNumber	int	`json:"PageNumber,omitempty"`
	SortBy		string	`json:"SortBy,omitempty"`
	OrderBy		string	`json:"OrderBy,omitempty"`
	TicketNumber	string  `json:"TicketNumber,omitempty"`
	CreatedTime	string  `json:"CreatedTime,omitempty"`
}


// This function is a handler that creates a POST request to update Ticket Common Agent Functions
//
// **Business Logic**: Function takes as an input a JSON Body and generates the response
//
// Returns data as shown in examples
func (h *Handler) PostGetTicketGridCustomer(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t GetTicketGridCustomer_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	utils.ResponseAbstract(postGetTicketGridCustomer(t),w)
}

func postGetTicketGridCustomer(T GetTicketGridCustomer_Request) []uint8{

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.postgetticketgridcustomer").GetString("uri")

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
