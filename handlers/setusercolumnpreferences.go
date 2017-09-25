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

type SUCP_Request struct {
	UserLogin      string   `json:"UserLogin"`
	Password       string   `json:"Password"`
	Action         string   `json:"Action"`
	ColumnSelected []string `json:"ColumnSelected"`
}

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
	//sessionIDString := callSessionDetails(T.UserLogin,T.Password)

	//url := felicitybaseurl+felicityapiuri+"?SessionID="+sessionIDString
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + T.UserLogin + "&Password=" + T.Password
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

	return bodyText

}
