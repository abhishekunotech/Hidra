package handlers

import (
<<<<<<< Updated upstream
	"github.com/Unotechsoftware/Hydra/utils"
=======
	"github.com/Unotechsoftware/Hydrav2/utils"
	"github.com/Unotechsoftware/Hydrav2/lerna"
>>>>>>> Stashed changes
	"net/http"
)

// This function is a handler that creates a GET API to get the count of closed tickets that were created/assigned for a customer user.
//
// **Business Logic**: Function takes as an input GET Parameter, __CustomerID__ that will identify Customer User and obtain a list of Tickets that are in "Closed" state.
//
// Returns data as shown in examples
func (h *Handler) GetCountofClosedTickets(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getcountofclosetickets"
        closedTicketCount := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
        utils.ResponseAbstract(closedTicketCount, w)

}
