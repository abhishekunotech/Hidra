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


type CustomerLanguageUpdate_Request struct{
	CustomerUserLogin	string	`json:"CustomerUserLogin,omitempty"`
	Password	string	`json:"Password,omitempty"`
	UserLanguage	string	`json:"UserLanguage,omitempty"`
    Group   string  `json:"Group,omitempty"`
}

// This function is a handler that creates a POST request to update Ticket Common Agent Functions
//
// **Business Logic**: Function takes as an input a JSON Body and generates the response
//
// Returns data as shown in examples
func (h *Handler) SetCustomerPreferencesLanguageUpdate(w http.ResponseWriter, r *http.Request) {
    fmt.Println("in lang up")
	decoder := json.NewDecoder(r.Body)
	var t CustomerLanguageUpdate_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	utils.ResponseAbstract(setCustomerPreferencesLanguageUpdate(t),w)
}

func setCustomerPreferencesLanguageUpdate(T CustomerLanguageUpdate_Request) []uint8{

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.setcustomerpreferenceslanguageupdate").GetString("uri")

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
