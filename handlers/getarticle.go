package handlers

import (
	"github.com/Unotechsoftware/Hydrav3/utils"
	"github.com/Unotechsoftware/Hydrav3/lerna"
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

// This function is a handler that creates a GET API that returns a List of Articles attached to a ticket
//
// **Business Logic**: Function creates a GET API that takes as an input, GET PARAMETERS : __TicketID__, __PageSize__ and __PageNumber__ to return a paginated JSON Response. The JSON Response returns the List of Articles attached to a Ticket identified by a TicketID
//
// Returns data as shown in examples
func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request) {

	mapHttp := utils.RequestAbstractGet(r)
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
