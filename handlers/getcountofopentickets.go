package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)


// This function is a handler that creates a GET API to get the count of open tickets that were created by/assigned to agent.
//
// **Business Logic**: Function that obtains a list of Tickets that are in "Open" state.
//
// Returns data as shown in examples
func (h *Handler) GetCountOfOpenTickets(w http.ResponseWriter, r *http.Request) {
	actionStrg := utils.RequestAbstractGet1(r)


	utils.ResponseAbstract(utils.ExecuteCallGet("components.otrs","components.otrs.apis.countofopentickets",actionStrg),w)
}
