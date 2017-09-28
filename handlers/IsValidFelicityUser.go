package handlers

import (
	"encoding/json"
	"github.com/antigloss/go/logger"
	"io"
	"io/ioutil"
	"net/http"
)

type LoginResult struct {
	LoginStatus   bool
	SessionString string
}

func checkValidUserDetails(username string, password string) (bool, string) {
	url := "http://192.168.2.166/felicity/nph-genericinterface.pl/Webservice/SessionAPI/SessionCreate?UserLogin=" + username + "&Password=" + password
	client := &http.Client{}
	var bodyReader io.Reader
	req, err := http.NewRequest("GET", url, bodyReader)
	if err != nil{
		logger.Error("Something went wrong, real bad",err.Error())
	}
	resp, err := client.Do(req)
	if err != nil{
		logger.Error("Something went wrong again",err.Error())
	}
	checkValidResult := true

	if err != nil {
		logger.Error("\n Session Creation Failed because - \n")
		logger.Error(err.Error())
		checkValidResult = false
		return checkValidResult, "nil"
	} else {
		req.Close = true
		bodyText, err := ioutil.ReadAll(resp.Body)
		var data SessionObject
		err = json.Unmarshal(bodyText, &data)

		if data.SessionIDStrg == "" {
			logger.Error("User Credentials Invalid")
			return false, "nil"
		}

		if err != nil {
			logger.Error("\n\n Json Unmarshaling failed\n\n")
			logger.Error(err.Error())
			return false, "nil"
		} else {
			return checkValidResult, data.SessionIDStrg
		}
	}
}

func (h *Handler) IsValidFelicityUser(w http.ResponseWriter, r *http.Request) {
	mapHttp := r.URL.Query()
	var userName string
	var password string
	for key, value := range mapHttp {
		if key == "username" {
			for _, valueStrg := range value {
				userName = valueStrg
			}
		}
		if key == "password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
	}

	validUser, sessionDataString := checkValidUserDetails(userName, password)

	var data LoginResult

	data.SessionString = sessionDataString
	data.LoginStatus = validUser

	jData, err := json.Marshal(data)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)


}
