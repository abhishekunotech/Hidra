package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"net/http"
	//"reflect"
	//"io"
)

func callListOfLinkedChange(w http.ResponseWriter, r *http.Request, username string, password string, ticketid string) {

	sessionIDString := callSessionDetails(username, password)

	fmt.Println("session id is ::", sessionIDString)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	fmt.Println("base url:- ", felicitybaseurl)
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlistoflinkedchange").GetString("uri")
	ticketid = ConfObj.Sub("components.otrs.apis.getlistoflinkedchange.parameters").GetString("TicketId")

	url := felicitybaseurl + felicityapiuri + "?TicketID=" + ticketid + "&SessionID=" + sessionIDString

	fmt.Println("url is::", url)
	res, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
	}

	bodyText, err := ioutil.ReadAll(res.Body)

	var data interface{}
	err = json.Unmarshal(bodyText, &data)
	if err != nil {
		logger.Error(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	/*json.NewEncoder(w).Encode(data)*/

}

//Function to get list of work orders
// Request as http://ip-host/getListOfWorkOrders?ticketID=521&password=abhik&userLogin=abhik

func (h *Handler) GetListOfLinkedChange(w http.ResponseWriter, r *http.Request) {
	//body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("in listof linked change")
	mapHttp := r.URL.Query()
	var userName string
	var password string
	var ticketid string
	for key, value := range mapHttp {
		if key == "ticketID" {
			for _, valueStrg := range value {
				ticketid = valueStrg
			}
		}
		if key == "userLogin" {
			for _, valueStrg := range value {
				userName = valueStrg
			}
		}
		if key == "password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
	}

	callListOfLinkedChange(w, r, userName, password, ticketid)

	//bodyStrg := string(body[:])
	//fmt.Fprintf(w,"www"+bodyStrg+"\n")
}
