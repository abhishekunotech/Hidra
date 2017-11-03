package handlers

import (
	"github.com/Unotechsoftware/Hydrav4/utils"
	"net/http"
)


// This function is a handler that creates a GET API that returns a List of Articles attached to a ticket
//
// **Business Logic**: Function creates a GET API that takes as an input, GET PARAMETERS : __TicketID__, __PageSize__ and __PageNumber__ to return a paginated JSON Response. The JSON Response returns the List of Articles attached to a Ticket identified by a TicketID
//
// Returns data as shown in examples
func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getarticle"
        article := utils.ExecuteCallGet(configStrg, uriStrg, actionStrg)
	utils.ResponseAbstract(utils.ExecuteCallGet("components.otrs","components.otrs.apis.getarticle",actionStrg),w)

}
