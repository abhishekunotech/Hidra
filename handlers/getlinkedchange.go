package handlers

import (
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io"
	"io/ioutil"
	"net/http"
)


//	url := "http://192.168.2.90:8080/felicity/nph-genericinterface.pl/Webservice/TicketAPI/ListOfLinkedChange?TicketID="+ticketid+"&SessionID="+sessionIDString
func callLinkedChanges(w http.ResponseWriter, r *http.Request, username string, password string, ticketid string) {

	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlinkedchanges").GetString("uri")
	felicityticketid := ConfObj.Sub("components.otrs.apis.getlinkedchanges.parameters").GetString("TicketID")
	url := felicitybaseurl + felicityapiuri + "?TicketID=" + felicityticketid + "&SessionID=" + sessionIDString
	client := &http.Client{}
	var bodyReader io.Reader
	req, err := http.NewRequest("GET", url, bodyReader)

	if err != nil {
		logger.Error("Get Request failed on call linked changes")
		logger.Error(err.Error())
	}
	resp, err := client.Do(req)

	if err != nil {
		logger.Error("Get Request Failed on call linked changes - Client do")
		logger.Error(err.Error())
	}

	if err != nil {
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

}

//Function to get the details about ticket.
// Request as http://ip-host/getLinkedChange?ticketID=627&password=abhik&username=abhik
func (h *Handler) GetLinkedChange(w http.ResponseWriter, r *http.Request) {
	mapHttp := r.URL.Query()
	var userName string
	var password string
	var ticketid string
	for key, value := range mapHttp {
		if key == "ticketID" {
			for _, valueStrg := range value {
				ticketid = valueStrg
			}
		}
		if key == "username" {
			for _, valueStrg := range value {
				userName = valueStrg
			}
		}
		if key == "password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
	}
	callLinkedChanges(w, r, userName, password, ticketid)
}
