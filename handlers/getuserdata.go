package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that provides the user data of requested User ID
// 
// **Business Logic**: Function takes Username, Password and User ID in Request Body to generate the response
//      
// Returns data as shown in examples
func (h *Handler) GetUserData(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.GetUserData"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)
}
