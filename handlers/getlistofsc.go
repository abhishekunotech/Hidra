package handlers

import (
	"github.com/Unotechsoftware/Hydrav2/lerna"
	"net/http"
	"github.com/Unotechsoftware/Hydrav2/utils"
)

func callGetCatalogList(username string, password string) []uint8{

	sessionIDString := callSessionDetails(username, password)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetCatalogList").GetString("uri")
	
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&SessionID=" + sessionIDString
	return utils.MakeHTTPGetCall(url)

}

// This function is a handler that creates a GET API to search for a term in the Service Catalog
//
// **Business Logic**: Function takes as an input GET Parameter, __term__ that will search for the value of that parameter within the Admin Catalog.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetCatalogList(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
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
			}

	something := callGetCatalogList(userName, password)
	utils.ResponseAbstract(something, w)
}
