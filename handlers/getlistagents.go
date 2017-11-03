package handlers

import (
<<<<<<< Updated upstream
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

=======
	"github.com/Unotechsoftware/Hydrav2/utils"
	"github.com/Unotechsoftware/Hydrav2/lerna"
	"net/http"
)

func callAgents(username string, password string, search string, term string) []uint8{

	sessionIDString := callSessionDetails(username, password)

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getlistagents").GetString("uri")

	url := felicitybaseurl + felicityapiuri + "?SessionID=" + sessionIDString + "&Search=" + search + "&Term=" + term
bodyText := utils.MakeHTTPGetCall(url)
	return bodyText
}

>>>>>>> Stashed changes
// This function is a handler that creates a GET API to search for an agent in the system
//
// **Business Logic**: Function takes as an input GET Parameter, __term__ that will search for agents whose login names match the parameters.
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetListAgents(w http.ResponseWriter, r *http.Request) {

	actionStrg := utils.RequestAbstractGet1(r)
        configStrg := "components.otrs"
        uriStrg := "components.otrs.apis.getlistagents"
        listAgents := utils.ExecuteCallGet(configStrg,uriStrg,actionStrg)
        utils.ResponseAbstract(listAgents, w)

}
