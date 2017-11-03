package handlers

import (
	"github.com/Unotechsoftware/Hydra/lerna"
        "net/http"
        "github.com/Unotechsoftware/Hydra/utils"
	"encoding/json"
	"github.com/antigloss/go/logger"
	"time"
	"io"
	"io/ioutil"
)

// Struct that defines the response of the IsUserLoggedIn API
type LoginResult struct {
	LoginStatus   bool
	SessionString string
}




func checkValidUserDetails(username string, password string) (bool, string) {


	ConfObj := lerna.ReturnConfigObject()
        felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
        felicityapiuri := ConfObj.Sub("components.otrs.apis.IsValidFelicityUser").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password
	client := &http.Client{}
	var bodyReader io.Reader

	start := time.Now()

	req, err := http.NewRequest("GET", url, bodyReader)
	if err != nil{
		logger.Error("Something went wrong, real bad",err.Error())
	}
	resp, err := client.Do(req)

	till := time.Since(start).String()

	logger.Info(" The call for the URL "+url+" took "+till+" to execute")

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


// Function that checks if the User is a Valid Felicity Service Desk User
//
// Business Logic: Checks if a session can be created for the credentials provided in the GET URL Parameters: __username__ and __password__ 
//
// Returns Status as false/true and if True Returns SessionID
func (h *Handler) IsValidFelicityUser(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
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

        
	validuser, session_data_strg := checkValidUserDetails(userName, password)
var data LoginResult

       data.SessionString = session_data_strg
       data.LoginStatus = validuser

       jData, err := json.Marshal(data)
       if err != nil {
               logger.Error(err.Error())
               return
       }
       w.Header().Set("Content-Type", "application/json")
       w.Write(jData)


}
