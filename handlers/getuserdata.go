package handlers

import (
	"github.com/Unotechsoftware/Hydrav2/utils"
	"github.com/Unotechsoftware/Hydrav2/lerna"
	"net/http"
)

func getuserdata(username string, password string, userid string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetUserData").GetString("uri")

	url := felicitybaseurl + felicityapiuri + "?UserID=" + userid + "&UserLogin=" + username +"&Password="+password

	return utils.MakeHTTPGetCall(url)

}

// This function is a handler that provides the user data of requested User ID
// 
// **Business Logic**: Function takes Username, Password and User ID in Request Body to generate the response
//      
// Returns data as shown in examples
func (h *Handler) GetUserData(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var userName string
	var password string
	var userid string
	for key, value := range mapHttp {
		if key == "UserID" {
			for _, valueStrg := range value {
				userid = valueStrg
			}
		}
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

	ciList := getuserdata(userName, password, userid)
	utils.ResponseAbstract(ciList,w)
}
