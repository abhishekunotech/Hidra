package handlers

import (
<<<<<<< Updated upstream
	"github.com/Unotechsoftware/Hydrav4/utils"
	"github.com/Unotechsoftware/Hydrav4/lerna"
=======
	"github.com/Unotechsoftware/Hydrav4/utils"
>>>>>>> Stashed changes
	"net/http"
)

// This function is a handler that provides the type of associated ID.
//
// **Business Logic**: Function takes Username and Password in Request Body to generate response.
//
// Returns data as shown in examples
func (h *Handler) ListType(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
	utils.ResponseAbstract(utils.ExecuteCallGet("components.otrs","components.otrs.apis.listtype",actionStrg),w)

}
