package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io"
	"io/ioutil"
	"net/http"
)

func callGetTicketHistoryVersionTwo(username string, password string, ticketid string, pagesize string, pagenumber string) []uint8{
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.gettickethistory").GetString("uri")
	felicityaction := ConfObj.Sub("components.otrs.apis.gettickethistory.parameters").GetString("Action")
	url := felicitybaseurl + felicityapiuri + "?TicketID=" + ticketid + "&UserLogin=" + username + "&Password=" + password + "&Action=" + felicityaction + "&PageSize=" + pagesize + "&PageNumber=" + pagenumber
	client := &http.Client{}
	var bodyReader io.Reader
	req, err := http.NewRequest("GET", url, bodyReader)
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("\n\nThis caused the following error \n\n")
		logger.Error(err.Error())
	}
	req.Close = true
	bodyText, err := ioutil.ReadAll(resp.Body)
	return bodyText
}
func (h *Handler) GetTicketHistoryVersionTwo(w http.ResponseWriter, r *http.Request) {

	mapHttp := r.URL.Query()
	var userName string
	var password string
	var ticketid string
	var PageSize string
	var PageNumber string
	for key, value := range mapHttp {
		if key == "UserLogin" {
			for _, valueStrg := range value {
				userName = valueStrg
			}
		}
		if key == "Password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
		if key == "TicketID" {
			for _, valueStrg := range value {
				ticketid = valueStrg
			}
		}
		if key == "PageSize" {
			for _, valueStrg := range value {
				PageSize = valueStrg
			}
		}
		if key == "PageNumber" {
			for _, valueStrg := range value {
				PageNumber = valueStrg
			}
		}
	}

	utils.ResponseAbstract(callGetTicketHistoryVersionTwo(userName, password, ticketid, PageSize, PageNumber),w)

}
