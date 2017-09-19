package handlers

import (
	"fmt"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io"
	"io/ioutil"
	"net/http"
)

func callGetPublicFAQ(w http.ResponseWriter, r *http.Request, username string, password string, itemid string) {

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetPublicFAQ").GetString("uri")
	sessionIDString := callSessionDetails(username, password)

	fmt.Println(felicityapiuri)
	fmt.Println("In public faqs")
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&ItemID=" + itemid + "&SessionID=" + sessionIDString
	fmt.Println("url",url)
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
	var data interface{}
	err = json.Unmarshal(bodyText, &data)
	if err != nil {
		logger.Error(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}


func (h *Handler) GetPublicFAQ(w http.ResponseWriter, r *http.Request) {
	//body, _ := ioutil.ReadAll(r.Body)
	mapHttp := r.URL.Query()
	var itemid string
	var username string
	var password string
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
		if key == "ItemID" {
                        for _, valueStrg := range value {
                                itemid = valueStrg
                        }
                }

	}
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(itemid)
	callGetPublicFAQ(w, r, username, password, itemid)

}
