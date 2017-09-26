package handlers

import (
	"fmt"
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"net/http"
)

type Close_Request struct{
	UserLogin	string	`json:"UserLogin"`
	Password	string	`json:"Password"`
	Action		string	`json:"Action"`
	TicketIDs	[]string	`json:"TicketIDs"`
	Subaction	string	`json:"Subaction"`
	ArticleType	string	`json:"ArticleType"`
	Unlock		string	`json:"Unlock"`
	StateID		string	`json:"StateID"`
	Subject		string	`json:"Subject"`
	Body		string	`json:"Body"`
}

func (h *Handler) PostCloseAgentTicket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IN close tik")
	decoder := json.NewDecoder(r.Body)
	var t Close_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()
	utils.ResponseAbstract(PostCloseAgentTicket(t),w)
}

func PostCloseAgentTicket(T Close_Request) []uint8 {

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.postcloseagentticket").GetString("uri")
	//sessionIDString := callSessionDetails(T.UserLogin,T.Password)

	//url := felicitybaseurl+felicityapiuri+"?SessionID="+sessionIDString
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

	return bodyText

}
