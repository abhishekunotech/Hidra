package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)


// This function is a handler that creates a GET API that returns the list of Configuration Items attached to a ticket
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ to identify a Ticket and return the list of Configuration Items assigned to the ticket.
//
// Returns data as shown in the examples.
func (h *Handler) GetListOfCIs(w http.ResponseWriter, r *http.Request) {
	actionStrng := utils.RequestAbstractGet1(r)
	configStrng := "components.otrs"
 	uriStrng := "components.otrs.apis.GetListOfCIs"
	listofCIs := utils.ExecuteCallGet(configStrng, uriStrng, actionStrng)
	utils.ResponseAbstract(listofCIs, w)
}
