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

// This function is a handler that creates a GET API to get the customer user details assigned to a ticket.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ that will identify Customer User and obtain personal details of the Customer.
//
// Returns data as shown in examples
func (h *Handler) GetCustomerInfobyTicketID(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getcustomerinfobyticketid"
        custInfo := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
        utils.ResponseAbstract(custInfo, w)

}
