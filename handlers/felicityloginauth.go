package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callFelicityLoginAuth(username string, password string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.FelicityLoginAuth").GetString("uri")
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password="+password
	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that creates a GET API that returns the UserID and SessionID.
//
// **Business Logic**: Function takes as an input userlogin and password to return UserID and SessionID
//
// Returns data as Error or {"UserID","SessionID"}
func (h *Handler) FelicityLoginAuth(w http.ResponseWriter, r *http.Request) {
	mapHttp := r.URL.Query()
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
	}
	utils.ResponseAbstract(callFelicityLoginAuth(username, password),w)

}
