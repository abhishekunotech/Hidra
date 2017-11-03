
package handlers

import (
<<<<<<< Updated upstream
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

=======
	"github.com/Unotechsoftware/Hydrav2/utils"
	"net/http"
)


>>>>>>> Stashed changes
// This function is a handler that creates a GET API that returns a list of dynamic fields attached to a Ticket and their corresponding values.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ identifies the Ticket and returns a list of Dynamic Field names and corresponding values.
//
// Returns data as shown in examples.
func (h *Handler) GetTicketDynamicFieldPossible(w http.ResponseWriter, r *http.Request) {
<<<<<<< Updated upstream

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.GetTicketDynamicFieldPossible"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)
=======
	//body, _ := ioutil.ReadAll(r.Body)
//	mapHttp := r.URL.Query()
	actionStrg := utils.RequestAbstractGet1(r)
	utils.ResponseAbstract(utils.ExecuteCallGet("components.otrs","components.otrs.apis.GetTicketDynamicFieldPossible",actionStrg),w)
>>>>>>> Stashed changes

}
