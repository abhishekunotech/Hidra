package handlers

import (
<<<<<<< Updated upstream
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
=======
	"github.com/Unotechsoftware/Hydra/utils"
>>>>>>> Stashed changes
	"net/http"
)


// This function is a handler that displays the priority with the associated ID.
//
// **Business Logic**: Function uses Username and Password in Request Body to generate response.
//
// Returns data as shown in examples
func (h *Handler) ListPriority(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
	utils.ResponseAbstract(utils.ExecuteCallGet("components.otrs","components.otrs.apis.listpriority",actionStrg),w)

}
