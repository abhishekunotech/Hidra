 package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
//	"reflect"
)

func callLinkedTickets(w http.ResponseWriter, r *http.Request, username string, password string, ticketid string){
	
	sessionIDString := callSessionDetails(username,password)
	fmt.Printf("Session ID in call linked Changes is "+sessionIDString)
	url := "http://192.168.2.90:8080/felicity/nph-genericinterface.pl/Webservice/TicketAPI/ListOfLinkedTickets?TicketID="+ticketid+"&SessionID="+sessionIDString
	client := &http.Client{}
	var bodyReader io.Reader
    	req, err := http.NewRequest("GET", url,bodyReader)
    	//req.SetBasicAuth(username,password)
    	//req.Header.Set("Authorization", "Basic Z2xwaTpnbHBp")
    	resp, err := client.Do(req)
//	req.Close = true
    	if err != nil{
		fmt.Printf("\n\nThis caused the following error \n\n")
        	fmt.Printf(err.Error())
    	}
	req.Close = true
    	bodyText, err := ioutil.ReadAll(resp.Body)
	var data interface{}
    	err = json.Unmarshal(bodyText, &data)
   	if err != nil {
        	fmt.Printf(err.Error())
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
