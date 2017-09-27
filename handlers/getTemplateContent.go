package handlers

import (
//	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"io/ioutil"
	"net/http"
	"github.com/Unotechsoftware/Hydra/utils"
)

func callContentTemplate(username string, password string, templateID string) []uint8 {

	sessionIDString := callSessionDetails(username, password)

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.templatecontent").GetString("uri")

	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString + "&TemplateID=" + templateID

	res := utils.MakeHTTPGetCall(url)

	bodyText, err := ioutil.ReadAll(res.Body)
	return bodyText

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
