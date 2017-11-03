package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

// This function is a handler that creates a GET API that returns the type of Articles
//
// **Business Logic**: Function takes as an input GET Parameter that returns the type of Article.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListArticleType(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.listarticletype"
        listarticle := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
        utils.ResponseAbstract(listarticle, w)

}
