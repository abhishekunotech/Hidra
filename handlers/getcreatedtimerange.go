package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API to get the Created Time Range.
//
// **Business Logic**: Function that returns the Created Time Range.
//
// Returns JSON Array of the format "id","minutes","value"; where value = Within x hours/minutes/days
func (h *Handler) GetCreatedTimeRange(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.GetCreatedTimeRange"
        timeRange := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
        utils.ResponseAbstract(timeRange, w)
}
