package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io"
	"io/ioutil"
	"net/http"
)

func callGetCategoryItemList(username string, password string, category string, id string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetCategoryItemList").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&Category=" + category + "&ID=" + id + "&SessionID=" + sessionIDString

	client := &http.Client{}
	var bodyReader io.Reader
	req, err := http.NewRequest("GET", url, bodyReader)

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("\n\nThis caused the following error \n\n")
		logger.Error(err.Error())
	}
	req.Close = true
	bodyText, err := ioutil.ReadAll(resp.Body)
	return bodyText
}


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
