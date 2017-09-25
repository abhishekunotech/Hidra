package handlers

import (
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"net/http"
	"github.com/Unotechsoftware/Hydra/utils"
)

func callContentTemplate(username string, password string, templateID string) interface {} {

	sessionIDString := callSessionDetails(username, password)

	logger.Info("session id is ::", sessionIDString)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	logger.Info("base url:- ", felicitybaseurl)
	felicityapiuri := ConfObj.Sub("components.otrs.apis.templatecontent").GetString("uri")

	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString + "&TemplateID=" + templateID

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

	return data

}

//Function to get list of work orders
// Request as http://ip-host/getTemplateContent?TemplateID=521&password=abhik&userLogin=abhik

func (h *Handler) GetTemplateContent(w http.ResponseWriter, r *http.Request) {

	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	var templateID string
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
		if key == "TemplateID" {
			for _, valueStrg := range value {
				templateID = valueStrg
			}
		}
	}

	interface1 := callContentTemplate(userName, password, templateID)
	utils.ResponseAbstract(interface1, w)
}
