 package handlers

import (
	"fmt"
	"encoding/json"
	"github.com/antigloss/go/logger"
	"net/http"
	"io/ioutil"
	"io"
	"github.com/Unotechsoftware/Hydra/lerna"
//	"reflect"
)

func callCustomerInfo(w http.ResponseWriter, r *http.Request, ticketid string){
	
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri :=  ConfObj.Sub("components.otrs.apis.getcustomerinfobyticketid").GetString("uri")
	felicityusername := ConfObj.Sub("components.otrs.apis.getcustomerinfobyticketid").GetString("userlogin")
	felicitypassword := ConfObj.Sub("components.otrs.apis.getcustomerinfobyticketid").GetString("password")
	sessionIDString := callSessionDetails(felicityusername,felicitypassword)
	url := felicitybaseurl+felicityapiuri+"?TicketID="+ticketid+"&SessionID="+sessionIDString
	fmt.Println("URL Meow")
	fmt.Println(url)
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
	w.Header().Set("Content-Type", "application/json")
    	json.NewEncoder(w).Encode(data)		
	/*json.NewEncoder(w).Encode(data)*/

}

//Function to get the details about ticket.

func GetCustomerInfobyTicketID(w http.ResponseWriter, r *http.Request) {
	//body, _ := ioutil.ReadAll(r.Body)
	mapHttp := r.URL.Query()
	var ticketid string
	for key,value := range mapHttp {
		if key == "ticketID"{
			for _, valueStrg := range value {
				ticketid = valueStrg
			}
		}
	}
	callCustomerInfo(w,r,ticketid)

}