package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API that returns the list of Changes linked to a ticket
//
// **Business Logic**: Function takes as an input GET Parameter, __ticketID__ that identifies a ticket and returns Linked Changes
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListOfLinkedChange(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getlistoflinkedchange"
	utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)

}
