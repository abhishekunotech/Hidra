package handlers

import (
	"github.com/antigloss/go/logger"
        "encoding/json"

        "net/http"
        "io/ioutil"
	//"reflect"
        //"io"
)
/*
type SessionObject struct{
	SessionIDStrg	string	`json:"SessionID"`
}

func callSessionDetails(username string, password string) string{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.SessionAPI").GetString("uri")
	felicityusername := ConfObj.Sub("components.otrs.apis.SessionAPI.parameters").GetString("UserLogin")
	felicitypassword := ConfObj.Sub("components.otrs.apis.SessionAPI.parameters").GetString("Password")
        url := felicitybaseurl+felicityapiuri+"?UserLogin="+felicityusername+"&Password="+felicitypassword
fmt.Println(url)
//	url := "http://192.168.2.90:8080/felicity/nph-genericinterface.pl/Webservice/SessionAPI/SessionCreate?UserLogin="+username+"&Password="+password
	
	client := &http.Client{}
	var bodyReader io.Reader
	req, err := http.NewRequest("GET", url, bodyReader)
	resp, err := client.Do(req)
	if err != nil{
		logger.Error("\n\n Session Creation failed because - \n\n")
		logger.Error(err.Error())
	}
	req.Close = true
	bodyText, err := ioutil.ReadAll(resp.Body)
	var data SessionObject
	err = json.Unmarshal(bodyText, &data)

	if err != nil{
		logger.Error(err.Error())
	}
	
	return data.SessionIDStrg	
}
*/
func callWorkOrders(w http.ResponseWriter, r *http.Request, username string, password string, ticketid string){

	sessionIDString := callSessionDetails(username,password)
        
	url := "http://192.168.2.90:8080/felicity/nph-genericinterface.pl/Webservice/TicketAPI/ListOfLinkedWorkorders?TicketID="+ticketid+"&SessionID="+sessionIDString

	res, err:= http.Get(url)
	if err != nil{
		logger.Error(err.Error())	
	}
       
	bodyText, err := ioutil.ReadAll(res.Body)
	
        var data interface{}
        err = json.Unmarshal(bodyText, &data)
        if err != nil {
                logger.Error(err.Error())
        }
        json.NewEncoder(w).Encode(data)
        /*json.NewEncoder(w).Encode(data)*/

}

//Function to get list of work orders
// Request as http://ip-host/getListOfWorkOrders?ticketID=521&password=abhik&userLogin=abhik

func GetListOfWorkOrders(w http.ResponseWriter, r *http.Request) {
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

        callWorkOrders(w,r,userName, password, ticketid)

        //bodyStrg := string(body[:])
        //fmt.Fprintf(w,"www"+bodyStrg+"\n")
}


