package handlers

import (
	"github.com/Unotechsoftware/Hydrav3/utils"
	"github.com/Unotechsoftware/Hydrav3/lerna"
	"net/http"
)

func callGetCategoryItemList(username string, password string, category string, id string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetCategoryItemList").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&Category=" + category + "&ID=" + id + "&SessionID=" + sessionIDString

	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that creates a GET API that returns a list of all the Items in a Service Catalog Category filtered by Item ID.
//
// **Business Logic**: Function takes as an input GET Parameter, __ID__ and __Category__ identify the Category in the Service Catalog and return a list of Items in that Category.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetCategoryItemList(w http.ResponseWriter, r *http.Request) {
	mapHttp := r.URL.Query()
	var id string
	var username string
	var password string
	var category string
	for key, value := range mapHttp {
		if key == "UserLogin" {
			for _, valueStrg := range value {
				username = valueStrg
			}
		}
		if key == "Password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
		if key == "Category" {
                        for _, valueStrg := range value {
                                category = valueStrg
                        }
                }
		if key == "ID" {
                        for _, valueStrg := range value {
                                id = valueStrg
                        }
                }

	}
	utils.ResponseAbstract(callGetCategoryItemList(username, password, category, id),w)

}
