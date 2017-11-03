package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API <TBD>
//
// **Business Logic**: Function takes as an input GET Parameter, __UserAccess__ <TBD>
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListGroupFilter(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
	configStrg := "components.otrs"
	uriStrg := "components.otrs.apis.getlistgroupfilter"
	groupfilter := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
	utils.ResponseAbstract(groupfilter, w)
}
