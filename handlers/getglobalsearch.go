package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)


// This function is a handler that creates a GET API to search for a term in the Service Catalog
//
// **Business Logic**: Function takes as an input GET Parameter, __term__ that will search for the value of that parameter within the Admin Catalog.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetGlobalSearch(w http.ResponseWriter, r *http.Request) {
	actionStrg := utils.RequestAbstractGet1(r)
	utils.ResponseAbstract(utils.ExecuteCallGet("components.otrs","components.otrs.apis.GetGlobalSearch", actionStrg),w)

}
