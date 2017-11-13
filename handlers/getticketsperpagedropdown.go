package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)


// This function is a handler that creates a GET API to get the pagination items count.
//
// **Business Logic**: Function takes as an input GET Parameter, __ticketID__ that will identify a ticket and return all the Items attached to it.
//
// Returns data as shown in examples.
func (h *Handler) GetTicketsPerPageDropdown(w http.ResponseWriter, r *http.Request) {
	actionStrg := utils.RequestAbstractGet1(r)
	configStrg := "components.otrs"
	uriStrg := "components.otrs.apis.GetTicketsPerPageDropdown"
	utils.ResponseAbstract(utils.ExecuteCallGet1(configStrg,uriStrg,actionStrg),w)

}
