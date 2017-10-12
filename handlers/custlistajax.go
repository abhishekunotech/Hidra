package handlers

import (
	"github.com/Unotechsoftware/Hydrav2/utils"
	"github.com/Unotechsoftware/Hydrav2/lerna"
	"net/http"
)

func callCustListAjax(username string, password string, search string, term string) []uint8{

	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.custlistajax").GetString("uri")

	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString + "&Search=" + search + "&Term=" + term

	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that shows the list of Customers as per filter criteria.
//
// **Business Logic**: Function exports a GET API that will accept GET Parameters __Search__ and __Term__ and return list of Customer Users that match the value of Term Parameter as *Term_Value*.
//
// Returns a JSON Body with a list of Customer Users.
func (h *Handler) CustListAjax(w http.ResponseWriter, r *http.Request) {

	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	var search string
	var term string
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
		if key == "Search" {
			for _, valueStrg := range value {
				search = valueStrg
			}
		}
		if key == "Term" {
			for _, valueStrg := range value {
				term = valueStrg
			}
		}
	}

	utils.ResponseAbstract(callCustListAjax(userName, password, search, term),w)

}
