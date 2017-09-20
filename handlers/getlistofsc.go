package handlers

import (
	"fmt"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"net/http"
)

func callGetCatalogList(w http.ResponseWriter, r *http.Request, username string, password string) {

	sessionIDString := callSessionDetails(username, password)
	fmt.Println("In Service Catalogue")
	logger.Info("session id is ::", sessionIDString)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	logger.Info("base url:- ", felicitybaseurl)
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetCatalogList").GetString("uri")
	
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&SessionID=" + sessionIDString
	fmt.Println("URL :", url)
	res, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
	}

	bodyText, err := ioutil.ReadAll(res.Body)

	var data interface{}
	err = json.Unmarshal(bodyText, &data)
	if err != nil {
		logger.Error(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}


func (h *Handler) GetCatalogList(w http.ResponseWriter, r *http.Request) {

	mapHttp := r.URL.Query()
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

	callGetCatalogList(w, r, userName, password)

}
