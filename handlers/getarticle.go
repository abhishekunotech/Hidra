package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callArticle(username string, password string, ticketid string, pagesize string, pagenumber string) []uint8{

	sessionIDString := callSessionDetails(username, password)

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getarticle").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString + "&TicketID=" + ticketid + "&PageSize=" + pagesize + "&PageNumber=" + pagenumber

	return utils.MakeHTTPGetCall(url)

}


func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request) {

	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	var ticketid string
	var PageSize string
	var PageNumber string
	for key, value := range mapHttp {
		if key == "userLogin" {
			for _, valueStrg := range value {
				userName = valueStrg
			}
		}
		if key == "password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
		if key == "ticketID" {
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

	utils.ResponseAbstract(callArticle(userName, password, ticketid, PageSize, PageNumber),w)

}
