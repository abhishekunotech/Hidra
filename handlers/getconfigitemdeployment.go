package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API to get Link Object CI Deployment State Dropdown List. 
//
// **Business Logic**: Function takes as an input GET Parameter UserLogin and Password.
//
// Returns data as shown in examples
func (h *Handler) GetConfigItemDeploymentState(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.GetConfigItemDeploymentState"
        configDeployment := utils.ExecuteCallGet(configStrg, uriStrg, actionStrg)
        utils.ResponseAbstract(configDeployment, w)
}
