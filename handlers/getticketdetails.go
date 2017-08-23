package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)


func callTicketDetails(w http.ResponseWriter, r *http.Request, username string, password string, ticketid string){
	//url := "http://192.168.2.181/felicity/nph-genericinterface.pl/Webservice/GetTicketDetails/GetTicketDetails/"+ticketid+"?UserLogin="+username+"&Password="+password
        url := "http://172.17.0.2:8080/getCIDetails"
	res, err := http.Get(url)

        //Errors are handled if any.
        if err != nil {
                panic(err.Error())
        }

        //ReadAll reads from response until an error or EOF and returns the data it read.
        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
                panic(err.Error())
        }
        var data interface{}

        //To decode JSON data, use the Unmarshal function.
        err = json.Unmarshal(body, &data)
        if err != nil {
                panic(err.Error())
        }
        fmt.Printf("Results: %v\n", data)
        //Encode the data
	//return data
	
        json.NewEncoder(w).Encode(data)	
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
	fmt.Fprintf(w,"meow")	
	callTicketDetails(w,r,userName, password, ticketid)

	//bodyStrg := string(body[:])
	//fmt.Fprintf(w,"www"+bodyStrg+"\n")
}
