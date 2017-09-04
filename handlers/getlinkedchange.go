package handlers

import (
	"fmt"
	"github.com/Unotechsoftware/Hydra/lerna"
	"encoding/json"
	"github.com/antigloss/go/logger"
	"net/http"
	"io/ioutil"
	"io"
)
/*
type SessionObject struct{
	SessionIDStrg	string	`json:"SessionID"`
}

func callSessionDetails(username string, password string) string{
	
	ConfObj := lerna.ReturnConfigObject()
        felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
        felicityapiuri := ConfObj.Sub("components.otrs.apis").GetString("uri")
        felicityusername := ConfObj.Sub("components.otrs.apis.SessionAPI.parameters").GetString("UserLogin")
        felicitypassword := ConfObj.Sub("components.otrs.apis.SessionAPI.parameters").GetString("Password")
        url := felicitybaseurl+felicityapiuri+"?UserLogin="+felicityusername+"&Password="+felicitypassword
	fmt.Println(url)	

//url := "http://192.168.2.90:8080/felicity/nph-genericinterface.pl/Webservice/SessionAPI/SessionCreate?UserLogin="+username+"&Password="+password
//	fmt.Printf("\n\n url is "+url+"\n\n")
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
*/

//	url := "http://192.168.2.90:8080/felicity/nph-genericinterface.pl/Webservice/TicketAPI/ListOfLinkedChange?TicketID="+ticketid+"&SessionID="+sessionIDString
func callLinkedChanges(w http.ResponseWriter, r *http.Request, username string, password string, ticketid string){
	
	sessionIDString := callSessionDetails(username,password)
	ConfObj := lerna.ReturnConfigObject()
        felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
        felicityapiuri := ConfObj.Sub("components.otrs.apis.getlinkedchanges").GetString("uri")
        //felicityusername := ConfObj.Sub("components.otrs.apis.getlinkedchanges.parameters").GetString("UserLogin")
        //felicitypassword := ConfObj.Sub("components.otrs.apis.getlinkedchanges.parameters").GetString("Password")
	felicityticketid := ConfObj.Sub("components.otrs.apis.getlinkedchanges.parameters").GetString("TicketID")
        url := felicitybaseurl+felicityapiuri+"?TicketID="+felicityticketid+"&SessionID="+sessionIDString
        fmt.Println(url)

	client := &http.Client{}
	var bodyReader io.Reader
    	req, err := http.NewRequest("GET", url,bodyReader)
    	
	if err != nil {

		fmt.Println("Get Request failed on call linked changes")
		fmt.Println(err.Error())
	}
	//req.SetBasicAuth(username,password)
    	//req.Header.Set("Authorization", "Basic Z2xwaTpnbHBp")
    	resp, err := client.Do(req)

	if err != nil{
		fmt.Println("Get Request Failed on call linked changes - Client do")
		fmt.Println(err.Error())
	}

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
