package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API to get the Tickets Linked to a Ticket.
//
// **Business Logic**: Function takes as an input GET Parameter, __ticketID__ that will identify a ticket and return all the Tickets attached to it.
//
// Returns data as shown in examples.
func (h *Handler) GetLinkedTickets(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getlinkedticketdetails"
        linkedTicket := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
        utils.ResponseAbstract(linkedTicket, w)

}
