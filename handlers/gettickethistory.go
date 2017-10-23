package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that provides the history about requested Ticket 
//
// **Business Logic**: Function uses the Ticket ID, Username and Password in Request Body to generate the response
//
// Returns data as shown in examples
func (h *Handler) GetTicketHistory(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.gettickethistory"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)

}
