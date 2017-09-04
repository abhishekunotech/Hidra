package handlers

import (
	"fmt"
	"github.com/Unotechsoftware/Hydra/lerna"
	"encoding/json"
	"github.com/antigloss/go/logger"
	"net/http"
	"io/ioutil"
	"io"
	//"reflect"
)


func callTicketDetails(w http.ResponseWriter, r *http.Request, username string, password string, ticketid string){
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getticketdetails").GetString("uri")
	felicityusername := ConfObj.Sub("components.otrs.apis.getticketdetails.parameters").GetString("UserLogin")
	felicitypassword := ConfObj.Sub("components.otrs.apis.getticketdetails.parameters").GetString("Password")
        url := felicitybaseurl+felicityapiuri+"?UserLogin="+felicityusername+"&Password="+felicitypassword
fmt.Println(url)	
//	url := "http://192.168.2.152/felicity/nph-genericinterface.pl/Webservice/TicketAPI/TicketGet/"+ticketid+"?UserLogin="+username+"&Password="+password
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
// Request as http://ip-host/getTicketDetails?ticketID=521&password=abhik&userLogin=abhik
func GetTicketDetails(w http.ResponseWriter, r *http.Request) {
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
		if key == "userLogin"{
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
	callTicketDetails(w,r,userName, password, ticketid)

	//bodyStrg := string(body[:])
	//fmt.Fprintf(w,"www"+bodyStrg+"\n")
}
