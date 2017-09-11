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
func callAgents(w http.ResponseWriter, r *http.Request, username string, password string, search string, term string){

	sessionIDString := callSessionDetails(username,password)

	fmt.Println("session id is ::",sessionIDString)        
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	fmt.Println("base url:- ",felicitybaseurl)
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlistagents").GetString("uri")
		
	url := felicitybaseurl+felicityapiuri+"?SessionID="+sessionIDString+"&Search="+search+"&Term="+term

	//fmt.Println("url is::",url)	
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
// Request as http://ip-host/getListOfWorkOrders?ticketID=521&password=abhik&userLogin=abhik

func (h *Handler) GetListAgents(w http.ResponseWriter, r *http.Request) {
        //body, _ := ioutil.ReadAll(r.Body)

	mapHttp := r.URL.Query()
        var userName string
        var password string
	var search string
	var term   string
        //var ticketid string
        for key,value := range mapHttp {
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
		if key == "Search"{
			for _, valueStrg := range value {
				search = valueStrg	
			}
		}
		if key == "Term"{
			for _, valueStrg := range value {
				term = valueStrg
			}
		}
        }

        callAgents(w,r,userName, password, search, term)

}


