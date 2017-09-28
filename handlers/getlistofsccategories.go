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
