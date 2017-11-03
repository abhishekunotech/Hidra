package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API that returns the details of the Ticket.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ identifies the Ticket and returns the details of that ticket.
//
// Returns data as shown in examples.
func (h *Handler) GetTicketDetails(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getticketdetails"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)
}
