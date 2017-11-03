package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API to search for an agent in the system
//
// **Business Logic**: Function takes as an input GET Parameter, __term__ that will search for agents whose login names match the parameters.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListAgents(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getlistagents"
        listAgents := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
        utils.ResponseAbstract(listAgents, w)

}
