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
type LinkObject_Request struct{
	UserLogin	string	`json:"UserLogin,omitempty"`
	Password	string	`json:"Password,omitempty"`
	SourceObject		string	`json:"SourceObject,omitempty"`
	LinkTargetKeys	[]string	`json:"LinkTargetKeys,omitempty"`
	SourceKey	string	`json:"SourceKey,omitempty"`
	TypeIdentifier	string	`json:"TypeIdentifier,omitempty"`
	TargetIdentifier	string	`json:"TargetIdentifier,omitempty"`
}

// This function is a handler that creates a POST API to link objects.
//
// **Business Logic**: Function takes as input a JSON Body and links objects based on the parameters in the JSON Request.
//
// Returns result message as success or failed.
func (h *Handler) LinkObject(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t LinkObject_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()
	utils.ResponseAbstract(LinkObject(t),w)
}

func LinkObject(T LinkObject_Request) []uint8 {

	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.LinkObject").GetString("uri")

	url := felicitybaseurl + felicityapiuri 
	j, err := json.Marshal(T)

	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)
	
	return utils.MakeHTTPPostCall(url,b)

}
