
package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)


// This function is a handler that creates a GET API that returns a list of dynamic fields attached to a Ticket and their corresponding values.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ identifies the Ticket and returns a list of Dynamic Field names and corresponding values.
//
// Returns data as shown in examples.
func (h *Handler) GetTicketDynamicFieldPossible(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.GetTicketDynamicFieldPossible"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)

}
