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

// This function is a handler that creates a GET API that returns the SLA information.
//
// **Business Logic**: Function takes as an input GET Parameter, __TicketID__ identifies the Ticket and returns the information of the SLA assigned to the ticket.
//
// Returns data as shown in examples.
func (h *Handler) GetSLAInfo(w http.ResponseWriter, r *http.Request) {
	actionStrng := utils.RequestAbstractGet1(r)
	configStrng := "components.otrs"
	uriStrng := "components.otrs.apis.getslainfo"
	slainfo := utils.ExecuteCallGet(configStrng, uriStrng, actionStrng)
	utils.ResponseAbstract(slainfo, w)
}
