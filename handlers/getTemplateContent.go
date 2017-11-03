package handlers

import (
	"net/http"
	"github.com/Unotechsoftware/Hydra/utils"
)

// This function is a handler that creates a GET API that returns the content of a Template defined by its ID.
//
// **Business Logic**: Function takes as an input GET Parameter, __TemplateID__ identifies the Template and returns the content of that template.
//
// Returns data as shown in examples.
func (h *Handler) GetTemplateContent(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.templatecontent"
        utils.ResponseAbstract(utils.ExecuteCallGet(configStrg, uriStrg, actionStrg), w)

}
