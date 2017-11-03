package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API that returns a List of Workorder Graphs assigned to a ticket.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ identifies the Ticket and returns a list of Workorders assigned to that ticket as a graph.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListofWorkorderGraph(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getlistofworkordergraph"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)
}
