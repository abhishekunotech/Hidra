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

func callCustListAjax(w http.ResponseWriter, r *http.Request, username string, password string, search string, term string) {

	sessionIDString := callSessionDetails(username, password)

	fmt.Println("session id is ::", sessionIDString)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	fmt.Println("base url:- ", felicitybaseurl)
	felicityapiuri := ConfObj.Sub("components.otrs.apis.custlistajax").GetString("uri")

	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString + "&Search=" + search + "&Term=" + term

	//fmt.Println("url is::",url)
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

func (h *Handler) CustListAjax(w http.ResponseWriter, r *http.Request) {
	//body, _ := ioutil.ReadAll(r.Body)

	mapHttp := r.URL.Query()
	var userName string
	var password string
	var search string
	var term string
	for key, value := range mapHttp {
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
		if key == "Search" {
			for _, valueStrg := range value {
				search = valueStrg
			}
		}
		if key == "Term" {
			for _, valueStrg := range value {
				term = valueStrg
			}
		}
	}

	callCustListAjax(w, r, userName, password, search, term)

}
