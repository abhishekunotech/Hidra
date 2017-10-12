package handlers

import (
	"github.com/Unotechsoftware/Hydrav3/utils"
	"github.com/Unotechsoftware/Hydrav3/lerna"
	"net/http"
)

func callUserColumnPreference(Component string,URI string ,Action string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	
	felicitybaseurl := ConfObj.Sub(Component).GetString("url")
	felicityapiuri := ConfObj.Sub(URI).GetString("uri")
	url := felicitybaseurl + felicityapiuri +  Action 
	return utils.MakeHTTPGetCall(url)
}

// This function is a handler that provides the details about user column preferences based on action 
//
// **Business Logic**: Function uses Username and Password in Request Body to generate the response
//
// Returns data as shown in examples
func (h *Handler) GetUserColumnPreferences(w http.ResponseWriter, r *http.Request) {
	actionStrg := utils.RequestAbstractGet1(r)
	componentStrg := "components.otrs"
	uriStrg := "components.otrs.apis.getusercolumnpreference"
	utils.ResponseAbstract(utils.ExecuteCallGet(componentStrg,uriStrg,actionStrg),w)

}
