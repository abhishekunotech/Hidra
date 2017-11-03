package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API to get linked object lists
//
// **Business Logic**: Function takes as an input GET Parameter UserLogin and password.
//
// Returns data as shown in examples
func (h *Handler) GetLinkObjectList(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.GetLinkObjectList"
        linkObject := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
        utils.ResponseAbstract(linkObject, w)

}
