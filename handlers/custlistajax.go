package handlers

import (
<<<<<<< Updated upstream
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
=======
	"github.com/Unotechsoftware/Hydrav2/utils"
>>>>>>> Stashed changes
	"net/http"
)


// This function is a handler that shows the list of Customers as per filter criteria.
//
// **Business Logic**: Function exports a GET API that will accept GET Parameters __Search__ and __Term__ and return list of Customer Users that match the value of Term Parameter as *Term_Value*.
//
// Returns a JSON Body with a list of Customer Users.
func (h *Handler) CustListAjax(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
	utils.ResponseAbstract(utils.ExecuteCallGet("components.otrs","components.otrs.apis.custlistajax",actionStrg),w)

}
