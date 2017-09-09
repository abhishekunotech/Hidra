package handlers

import (
	"encoding/json"
	"github.com/antigloss/go/logger"
	"net/http"
	"io/ioutil"
	"io"
	"fmt"
	"github.com/Unotechsoftware/Hydra/lerna"
)

func callGetTicketInfo(w http.ResponseWriter, r *http.Request, ticketid string, username string, password string){
	
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri :=  ConfObj.Sub("components.otrs.apis.GetTicketInfo").GetString("uri")
	sessionIDString := callSessionDetails(username,password)
	url := felicitybaseurl+felicityapiuri+"?TicketID="+ticketid+"&SessionID="+sessionIDString
	fmt.Println("URL is")
	fmt.Println(url)
	client := &http.Client{}
	var bodyReader io.Reader
    	req, err := http.NewRequest("GET", url,bodyReader)

    	resp, err := client.Do(req)
    	if err != nil{
		//logger.Error("\n\nThis caused the following error \n\n")
        	//logger.Error(err.Error())
		fmt.Println("\n\n This casued the following error in Request")
		fmt.Println(err.Error())
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

}

//Function to get the details about ticket.

func GetTicketInfo(w http.ResponseWriter, r *http.Request) {
	//body, _ := ioutil.ReadAll(r.Body)
	mapHttp := r.URL.Query()
	var ticketid string
	var username string
	var password string
	for key,value := range mapHttp {
		if key == "TicketID"{
			for _, valueStrg := range value {
				ticketid = valueStrg
			}
		}
		if key == "UserLogin"{
			for _, valueStrg := range value{
				username = valueStrg	
			}
		}
		if key == "Password"{
			for _, valueStrg := range value{
				password = valueStrg
			}
		}
	}
	fmt.Println("User name and Password is")
	fmt.Println(username)
	fmt.Println(password)
	callGetTicketInfo(w,r,ticketid,username,password)

}
