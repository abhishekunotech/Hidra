package handlers

import (
        "github.com/antigloss/go/logger"
        "encoding/json"
        "github.com/Unotechsoftware/Hydra/lerna"
        "net/http"
        "fmt"
        "io/ioutil"
        //"reflect"
        //"io"
)
func callTicketAll(w http.ResponseWriter, r *http.Request, username string, password string, ticketid string){

        sessionIDString := callCustomerSessionDetails(username,password)

        fmt.Println("session id is ::",sessionIDString)
        ConfObj := lerna.ReturnConfigObject()
        felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
        fmt.Println("base url:- ",felicitybaseurl)
        felicityapiuri := ConfObj.Sub("components.otrs.apis.TicketAll").GetString("uri")
        filter := ConfObj.Sub("components.otrs.apis.TicketAll.parameters").GetString("filter")


	
        url := felicitybaseurl+felicityapiuri+"?Filter="+filter+"&SessionID="+sessionIDString

        fmt.Println("url is::",url)
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
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
        /*json.NewEncoder(w).Encode(data)*/

}

//Function to get list of work orders
// Request as http://ip-host/ticketAll?ticketID=521&password=abhik&userLogin=abhik

func (h *Handler) TicketAll(w http.ResponseWriter, r *http.Request) {
        //body, _ := ioutil.ReadAll(r.Body)
	
	fmt.Println("in ticketall")
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

        callTicketAll(w,r,userName, password, ticketid)

        //bodyStrg := string(body[:])
        //fmt.Fprintf(w,"www"+bodyStrg+"\n")
}


