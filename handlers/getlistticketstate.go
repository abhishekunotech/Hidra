package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)


// This function is a handler that creates a GET API that returns a list of tickets and their corresponding states.
//
// **Business Logic**: Function returns a list of Tickets and their corresponding states.
//
// Returns data as shown in examples.
func (h *Handler) GetListTicketState(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getlistticketstate"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)
}
