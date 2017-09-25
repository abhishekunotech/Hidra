package handlers

import (
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"net/http"
	"github.com/Unotechsoftware/Hydra/utils"
)

func callGetCatalogList(username string, password string) []uint8{

	sessionIDString := callSessionDetails(username, password)
	logger.Info("session id is ::", sessionIDString)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	logger.Info("base url:- ", felicitybaseurl)
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetCatalogList").GetString("uri")
	
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&SessionID=" + sessionIDString
	res, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
	}

	bodyText, err := ioutil.ReadAll(res.Body)
	
	return bodyText

}


func (h *Handler) GetCatalogList(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
//	mapHttp := r.URL.Query()
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
