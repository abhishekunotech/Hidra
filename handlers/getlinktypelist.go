package handlers

import (
<<<<<<< Updated upstream
	"github.com/Unotechsoftware/Hydra/utils"
=======
	"github.com/Unotechsoftware/Hydra/utils"
>>>>>>> Stashed changes
	"net/http"
)


// This function is a handler that creates a GET API to get details about Link Type List. 
//
// **Business Logic**: Function takes as an input GET Parameters UserLogin, Password, SourceObject and TargetIdentifier and generate the response
//
// Returns data as shown in examples
func (h *Handler) GetLinkTypeList(w http.ResponseWriter, r *http.Request) {
	actionStrg := utils.RequestAbstractGet1(r)
	utils.ResponseAbstract(utils.ExecuteCallGet("components.otrs","components.otrs.apis.GetLinkTypeList",actionStrg),w)
}
