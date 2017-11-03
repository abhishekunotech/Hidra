package handlers

import (
<<<<<<< Updated upstream
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

=======
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)


>>>>>>> Stashed changes
// This function is a handler
//
// **Business Logic**: To be done.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListAssignedQueue(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
<<<<<<< Updated upstream
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getlistassignedqueue"
        listass := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
        utils.ResponseAbstract(listass, w)

=======
	assignedqueues := utils.ExecuteCallGet("components.otrs","components.otrs.apis.getlistassignedqueue",actionStrg)
	utils.ResponseAbstract(assignedqueues, w)
>>>>>>> Stashed changes
}
