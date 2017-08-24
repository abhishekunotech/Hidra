 package handlers

import (
	"encoding/json"
	"github.com/antigloss/go/logger"
	"net/http"
	"io/ioutil"
	"io"
//	"reflect"
)

func callLinkedTickets(w http.ResponseWriter, r *http.Request, username string, password string, ticketid string){
	
	sessionIDString := callSessionDetails(username,password)
	
	url := "http://192.168.2.901:8080/felicity/nph-genericinterface.pl/Webservice/TicketAPI/ListOfLinkedTickets?TicketID="+ticketid+"&SessionID="+sessionIDString
	client := &http.Client{}
	var bodyReader io.Reader
    	req, err := http.NewRequest("GET", url,bodyReader)
    	//req.SetBasicAuth(username,password)
    	//req.Header.Set("Authorization", "Basic Z2xwaTpnbHBp")
    	resp, err := client.Do(req)
//	req.Close = true
    	if err != nil{
		logger.Error("\n\nThis caused the following error \n\n")
        	logger.Error(err.Error())
    	}
	req.Close = true
    	bodyText, err := ioutil.ReadAll(resp.Body)
	var data interface{}
    	err = json.Unmarshal(bodyText, &data)
   	if err != nil {
        	logger.Error(err.Error())
    	}
    	json.NewEncoder(w).Encode(data)		
	/*json.NewEncoder(w).Encode(data)*/

}

//Function to get the details about ticket.

func GetLinkedTickets(w http.ResponseWriter, r *http.Request) {
	//body, _ := ioutil.ReadAll(r.Body)
	mapHttp := r.URL.Query()
	var userName string
	var password string
	var ticketid string
	for key,value := range mapHttp {
		if key == "ticketID"{
			for _, valueStrg := range value {
				ticketid = valueStrg
			}
		}
		if key == "username"{
			for _, valueStrg := range value {
				userName = valueStrg
			}
		}
		if key == "password"{
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
	}
	callLinkedTickets(w,r,userName, password, ticketid)

}
