package handlers

import (
	"encoding/json"
	"fmt"
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
	url := "http://192.168.2.181/felicity/nph-genericinterface.pl/Webservice/SessionAPI/SessionCreate?UserLogin=" + username + "&Password=" + password

	client := &http.Client{}
	var bodyReader io.Reader
	fmt.Println("URL is ")
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, bodyReader)

	resp, err := client.Do(req)

	checkValidResult := true

	if err != nil {
		logger.Error("\n Session Creation Failed because - \n")
		logger.Error(err.Error())
		fmt.Println("\n\n Session Creation Failed \n\n")
		fmt.Println(err.Error())
		checkValidResult = false
		return checkValidResult, "nil"
	} else {
		req.Close = true

		bodyText, err := ioutil.ReadAll(resp.Body)

		var data SessionObject
		fmt.Println("\n\n Data is \n\n")
		fmt.Println(bodyText)
		err = json.Unmarshal(bodyText, &data)

		if data.SessionIDStrg == "" {
			logger.Error("User Credentials Invalid")
			fmt.Println("User Credentials Invalid")
			return false, "nil"
		}

		if err != nil {
			fmt.Println("\n\n Json Unmarshaling failed\n\n")
			logger.Error(err.Error())
			fmt.Println(err.Error())
			return false, "nil"
		} else {
			return checkValidResult, data.SessionIDStrg
		}
	}
}

func (h *Handler) IsValidFelicityUser(w http.ResponseWriter, r *http.Request) {
	//body, _ := ioutil.ReadAll(r.Body)
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
		panic(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

	//	w.Header().Set("Content-Type", "application/json")

	//      json.NewEncoder(w).Encode(data)

}
