package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)


// This function is a handler
//
// **Business Logic**: To be done.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListAssignedQueue(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getlistassignedqueue"
        listass := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
        utils.ResponseAbstract(listass, w)
}
