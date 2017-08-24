package handlers

import (
        "encoding/json"
        "fmt"
        "net/http"
        "io/ioutil"
	"reflect"
        "io"
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

func callWorkOrders(w http.ResponseWriter, r *http.Request, username string, password string, ticketid string){

	sessionIDString := callSessionDetails(username,password)
        
	url := "http://192.168.2.90:8080/felicity/nph-genericinterface.pl/Webservice/TicketAPI/ListOfLinkedWorkorders?TicketID="+ticketid+"&SessionID="+sessionIDString

	res, err:= http.Get(url)
	if err != nil{
		fmt.Printf(err.Error())	
	}
       
	bodyText, err := ioutil.ReadAll(res.Body)
	
        var data interface{}
        err = json.Unmarshal(bodyText, &data)
        if err != nil {
                fmt.Printf(err.Error())
        }
        json.NewEncoder(w).Encode(data)
        /*json.NewEncoder(w).Encode(data)*/

}

//Function to get list of work orders
// Request as http://ip-host/getListOfWorkOrders?ticketID=521&password=abhik&userLogin=abhik

func GetListOfWorkOrders(w http.ResponseWriter, r *http.Request) {
        //body, _ := ioutil.ReadAll(r.Body)
        fmt.Printf("in work orders")
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
	fmt.Printf("in work orders , username is "+userName)
        callWorkOrders(w,r,userName, password, ticketid)

        //bodyStrg := string(body[:])
        //fmt.Fprintf(w,"www"+bodyStrg+"\n")
}


