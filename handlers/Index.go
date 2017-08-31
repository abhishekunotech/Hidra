package handlers

import (
	"net/http"
	//"encoding/json"
//	"io"
//	"io/ioutil"
//	"github.com/antigloss/go/logger"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "hello bla")
        /*
	var ticketid string
        var username string
        var password string
        ticketid = "627"
        username = "amol"
        password = "amol"

        url := "http://192.168.2.152/felicity/nph-genericinterface.pl/Webservice/TicketAPI/TicketGet/"+ticketid+"?UserLogin="+username+"&Password="+password

        client := &http.Client{}
        var bodyReader io.Reader
        req, err := http.NewRequest("GET", url,bodyReader)
        //req.SetBasicAuth(username,password)
        //req.Header.Set("Authorization", "Basic Z2xwaTpnbHBp")
        resp, err := client.Do(req)
//      req.Close = true
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
        }*/
}

