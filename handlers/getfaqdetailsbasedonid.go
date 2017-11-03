package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API to get the details of Public FAQ from its ID.
//
// **Business Logic**: Function takes as an input GET Parameter, __ItemID__ that will identify a public FAQ and get details of the FAQ.
//
// Returns data as shown in examples
func (h *Handler) GetPublicFAQ(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.GetPublicFAQ"
        publicfaq := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
        utils.ResponseAbstract(publicfaq, w)
}
