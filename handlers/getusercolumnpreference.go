package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)


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
