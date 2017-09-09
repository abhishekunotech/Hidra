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

func callCountOfOpenTicketsCustomerUser(w http.ResponseWriter, r *http.Request, custID string, username string, password string, custuser string) {

	sessionIDString := callSessionDetails(username, password)

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.CountOfOpenTicketsUser").GetString("uri")
	state := ConfObj.Sub("components.otrs.apis.CountOfOpenTicketsUser.parameters").GetString("state")
	url := felicitybaseurl + felicityapiuri + "?State=" + state + "&SessionID=" + sessionIDString + "&CustomerID=" + custID + "&CustomerUser=" + custuser

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

func GetCountOfOpenTicketsCustomerUser(w http.ResponseWriter, r *http.Request) {
	//body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("in openticket count")
	mapHttp := r.URL.Query()

	var userName string
	var password string
	var custID string
	var custUser string
	for key, value := range mapHttp {
		if key == "CustomerID" {
			for _, valueStrg := range value {
				custID = valueStrg
			}
		}
		if key == "UserLogin" {
			for _, valueStrg := range value {
				userName = valueStrg
			}
		}
		if key == "Password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
		if key == "CustomerUser" {
			for _, valueStrg := range value {
				custUser = valueStrg
			}

		}
	}

		callCountOfOpenTicketsCustomerUser(w, r, custID, userName, password, custUser)
		//bodyStrg := string(body[:])
		//fmt.Fprintf(w,"www"+bodyStrg+"\n")

	
}
