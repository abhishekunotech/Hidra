package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)
// This Type defines the Input JSON Body for POST request to set user column preferences.
//
// The structure consists of parameters such as Userlogin, Password, Action and selected Column
type SUCP_Request struct {
	UserLogin      string   `json:"UserLogin"`
	Password       string   `json:"Password"`
	Action         string   `json:"Action"`
	ColumnSelected []string `json:"ColumnSelected"`
}
// This function is a handler that creates a POST Request to set user column preferences.
//
// **Business Logic**: Function takes as an input a JSON Body and uses the parameters in Request Body to generate response.
//
// Returns data as shown in examples
func (h *Handler) SetUserColumnPreferences(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t SUCP_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	intermediate := setUserColumnPreferences(t)
	utils.ResponseAbstract(intermediate,w)
}

func setUserColumnPreferences(T SUCP_Request) []uint8 {

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.setusercolumnpreferences").GetString("uri")

	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + T.UserLogin + "&Password=" + T.Password
	j, err := json.Marshal(T)

	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)

	
	return utils.MakeHTTPPostCall(url,b)

}
