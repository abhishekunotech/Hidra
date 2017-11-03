package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API to get details about navigation bar of an agent on the basis of access.
//
// **Business Logic**: Function takes as an input GET Parameter UserLogin and Password that will identify the agent and obtain details of nav bar.
//
// Returns data as shown in examples
func (h *Handler) GetAgentNavBar(w http.ResponseWriter, r *http.Request) {

	actionStrng := utils.RequestAbstractGet1(r)
	configStrng := "components.otrs"
	uriStrng := "components.otrs.apis.GetAgentNavBar"
	agentNavBar := utils.ExecuteCallGet(configStrng, uriStrng, actionStrng)
	utils.ResponseAbstract(agentNavBar, w)
}
