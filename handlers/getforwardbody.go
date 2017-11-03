package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)



// This function is a handler that creates a GET API to get the Forward Body of Ticket.
//
// **Business Logic**: Function takes as an input GET Parameter, __ticketID__ that will identify a ticket and return all the Changes attached to it.
//
// Returns data as shown in examples.
func (h *Handler) GetForwardBody(w http.ResponseWriter, r *http.Request) {
	ActionStrg := utils.RequestAbstractGet1(r)
	utils.ResponseAbstract(utils.ExecuteCallGet("components.otrs","components.otrs.apis.getforwardbody",ActionStrg),w)
}
