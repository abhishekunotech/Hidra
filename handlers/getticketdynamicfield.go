
package handlers

import (
<<<<<<< Updated upstream
	"github.com/Unotechsoftware/Hydrav4/utils"
=======
	"github.com/Unotechsoftware/Hydrav4/utils"
	"github.com/Unotechsoftware/Hydrav4/lerna"
>>>>>>> Stashed changes
	"net/http"
)

// This function is a handler that creates a GET API that returns a list of dynamic fields attached to a Ticket and their corresponding values.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ identifies the Ticket and returns a list of Dynamic Field names and corresponding values.
//
// Returns data as shown in examples.
func (h *Handler) GetTicketDynamicField(w http.ResponseWriter, r *http.Request) {
<<<<<<< Updated upstream

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.GetTicketDynamicField"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)
=======
	//body, _ := ioutil.ReadAll(r.Body)
	actionStrg := utils.RequestAbstractGet1(r)
	utils.ResponseAbstract(utils.ExecuteCallGet("components.otrs","components.otrs.apis.GetTicketDynamicField",actionStrg),w)
>>>>>>> Stashed changes

}
