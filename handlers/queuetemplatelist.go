package handlers

import (
	"github.com/antigloss/go/logger"
        "encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
        "net/http"
        "io/ioutil"
)

func callQueueTemplateList(w http.ResponseWriter, r *http.Request, username string, password string, queueid string){

	sessionIDString := callSessionDetails(username,password)

	logger.Info("session id is ::",sessionIDString)        
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	logger.Info("base url:- ",felicitybaseurl)
	felicityapiuri := ConfObj.Sub("components.otrs.apis.queuetemplatelist").GetString("uri")
		
	url := felicitybaseurl+felicityapiuri+"?SessionID="+sessionIDString+"&QueueID="+queueid

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

func (h *Handler) GetQueueTemplateList(w http.ResponseWriter, r *http.Request) {
        //body, _ := ioutil.ReadAll(r.Body)

	mapHttp := r.URL.Query()
        var userName string
        var password string
	var queueid string
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
		if key == "QueueID"{
			for _, valueStrg := range value {
				queueid = valueStrg
			}
		}
        }

        callQueueTemplateList(w,r,userName, password, queueid)

}


