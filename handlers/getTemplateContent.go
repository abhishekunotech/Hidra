package handlers

import (
	"github.com/Unotechsoftware/Hydrav2/lerna"
	"net/http"
	"github.com/Unotechsoftware/Hydrav2/utils"
)

func callContentTemplate(username string, password string, templateID string) []uint8 {

	sessionIDString := callSessionDetails(username, password)

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.templatecontent").GetString("uri")

	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString + "&TemplateID=" + templateID

	bodyText := utils.MakeHTTPGetCall(url)

	return bodyText

}

// This function is a handler that creates a GET API that returns the content of a Template defined by its ID.
//
// **Business Logic**: Function takes as an input GET Parameter, __TemplateID__ identifies the Template and returns the content of that template.
//
// Returns data as shown in examples.
func (h *Handler) GetTemplateContent(w http.ResponseWriter, r *http.Request) {

	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	var templateID string
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
		if key == "TemplateID" {
			for _, valueStrg := range value {
				templateID = valueStrg
			}
		}
	}

	interface1 := callContentTemplate(userName, password, templateID)
	utils.ResponseAbstract(interface1, w)
}
