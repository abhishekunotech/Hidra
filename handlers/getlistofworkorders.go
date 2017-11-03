package handlers

import (
<<<<<<< Updated upstream
	"github.com/Unotechsoftware/Hydra/utils"
=======
	"github.com/Unotechsoftware/Hydrav2/utils"
	"github.com/Unotechsoftware/Hydrav2/lerna"
>>>>>>> Stashed changes
	"net/http"
)

// This function is a handler that creates a GET API that returns a list of Workorders assigned to a ticket.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ identifies the Ticket and returns a list of Workorders assigned to that ticket.
//
// Returns data as shown in examples.
func (h *Handler) GetListOfWorkOrders(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getlistofworkorders"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)
}
