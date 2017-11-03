package handlers

import (
	"net/http"
	"github.com/Unotechsoftware/Hydra/utils"
)


// This function is a handler that creates a GET API that returns a list of all the Service Catalog Categories filtered by Catalog ID.
//
// **Business Logic**: Function takes as an input GET Parameter, __ID__ and __Catalog__ identify the Catalog in the Service Catalog and return a list of Categories in that Catalog.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetCategoryList(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.GetCategoryList"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)

}
