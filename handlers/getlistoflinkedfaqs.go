package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API that returns a list of Linked FAQs as related to a Ticket.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ that will identify a Ticket and return a list of Linked FAQs.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListOfLinkedFAQs(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.GetListOfFAQs"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)

}
