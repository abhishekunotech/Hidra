package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)


type CustomerRefreshTime_Request struct{
	CustomerUserLogin	string	`json:"CustomerUserLogin,omitempty"`
	Password	string	`json:"Password,omitempty"`
	UserRefreshTime	int	`json:"UserRefreshTime,omitempty"`
    Group   string  `json:"Group,omitempty"`
}

// This function is a handler that creates a POST request to update Ticket Common Agent Functions
//
// **Business Logic**: Function takes as an input a JSON Body and generates the response
//
// Returns data as shown in examples
func (h *Handler) SetCustomerPreferencesRefreshTime(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t CustomerRefreshTime_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	utils.ResponseAbstract(setCustomerPreferencesRefreshTime(t),w)
}

func setCustomerPreferencesRefreshTime(T CustomerRefreshTime_Request) []uint8{

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.setcustomerpreferencesrefreshtime").GetString("uri")

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
