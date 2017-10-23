package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)


// This function is a handler that creates a GET API that returns a list of all the Items in a Service Catalog Category filtered by Item ID.
//
// **Business Logic**: Function takes as an input GET Parameter, __ID__ and __Category__ identify the Category in the Service Catalog and return a list of Items in that Category.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetCategoryItemList(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.GetCategoryItemList"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)

}
