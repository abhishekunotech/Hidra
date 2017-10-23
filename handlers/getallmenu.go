package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API to get all menu details of navigation bar.
//
// **Business Logic**: Function takes as an input GET Parameter UserLogin and Password that will identify the agent and obtain details of nav bar.
//
// Returns data as shown in examples
func (h *Handler) GetAllMenu(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.GetAllMenu"
        allMenu := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
        utils.ResponseAbstract(allMenu, w)	
}
