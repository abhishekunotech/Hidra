package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"reflect"
)

type SessionObject struct{
	SessionIDStrg	string	`json:"SessionID"`
}

func callSessionDetails(username string, password string) string{
	url := "http://192.168.2.90:8080/felicity/nph-genericinterface.pl/Webservice/SessionAPI/SessionCreate?UserLogin="+username+"&Password="+password
	fmt.Printf("\n\n url is "+url+"\n\n")
	client := &http.Client{}
	var bodyReader io.Reader
	req, err := http.NewRequest("GET", url, bodyReader)
	resp, err := client.Do(req)
	if err != nil{
		fmt.Printf("\n\n Session Creation failed because - \n\n")
		fmt.Printf(err.Error())
	}
	req.Close = true
	bodyText, err := ioutil.ReadAll(resp.Body)
	var data SessionObject
	err = json.Unmarshal(bodyText, &data)
	fmt.Printf("\n\n body is ")
	fmt.Printf(reflect.TypeOf(data).String())
	fmt.Printf("\n\n")
	if err != nil{
		fmt.Printf("\n\n Error Occured in unmarshalling Session JSON \n\n")
		fmt.Printf(err.Error())
	}
	fmt.Printf("\n\nSESSION ID IN CALL SESSION DETAILS IS "+data.SessionIDStrg+"\n\n")
	return data.SessionIDStrg	
}


func callLinkedChanges(w http.ResponseWriter, r *http.Request, username string, password string, ticketid string){
	
	sessionIDString := callSessionDetails(username,password)
	fmt.Printf("Session ID in call linked Changes is "+sessionIDString)
	url := "http://192.168.2.90:8080/felicity/nph-genericinterface.pl/Webservice/TicketAPI/ListOfLinkedChange?TicketID="+ticketid+"&SessionID="+sessionIDString
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
// Request as http://ip-host/getLinkedChange?ticketID=627&password=abhik&username=abhik
func GetLinkedChange(w http.ResponseWriter, r *http.Request) {
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
	callLinkedChanges(w,r,userName, password, ticketid)

	//bodyStrg := string(body[:])
	//fmt.Fprintf(w,"www"+bodyStrg+"\n")
}
