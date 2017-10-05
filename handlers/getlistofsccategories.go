package handlers

import (
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
	"github.com/Unotechsoftware/Hydra/utils"
)

func callGetCategoryList(username string, password string, catalog string, id string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetCategoryList").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&Catalog=" + catalog + "&ID=" + id + "&SessionID=" + sessionIDString

	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that creates a GET API that returns a list of all the Service Catalog Categories filtered by Catalog ID.
//
// **Business Logic**: Function takes as an input GET Parameter, __ID__ and __Catalog__ identify the Catalog in the Service Catalog and return a list of Categories in that Catalog.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetCategoryList(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var id string
	var username string
	var password string
	var catalog string
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
		if key == "Catalog" {
                        for _, valueStrg := range value {
                                catalog = valueStrg
                        }
                }
		if key == "ID" {
                        for _, valueStrg := range value {
                                id = valueStrg
                        }
                }

	}
	categorylist := callGetCategoryList(username, password, catalog, id)
	utils.ResponseAbstract(categorylist, w)
}
