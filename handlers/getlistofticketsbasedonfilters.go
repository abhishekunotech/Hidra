package handlers



import (
	"github.com/Unotechsoftware/Hydra/utils"
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)

// Defines a struct that can be used to store the values obtained from Felicity Service Desk to show a grid of tickets
//
// Variables and their corresponding JSON keys are defined hereby
type Grid_API struct {
	UserLogin         string   `json:"UserLogin,omitempty"`
	Password          string   `json:"Password,omitempty"`
	PageSize          int      `json:"PageSize,omitempty"`
	PageNumber        int      `json:"pageNumber,omitempty"`
	Action            string   `json:"Action,omitempty"`
	TicketNumber      string   `json:"TicketNumber,omitempty"`
	TypeID            int      `json:"TypeID,omitempty"`
	QueueIDs          []int    `json:"QueueIDs,omitempty"`
	StateIDs          []int    `json:"StateIDs,omitempty"`
	CustomerUserLogin []string `json:"CustomerUserLogin,omitempty"`
	PriorityIDs       []int    `json:"PriorityIDs,omitempty"`
	CreatedUserIDs    []int    `json:"CreatedUserIDs,omitempty"`
	SourceIDs         []int    `json:"SourceIDs,omitempty"`
	Due               string   `json:"Due,omitempty"`
	GroupIDs          []string    `json:"GroupIDs,omitempty"`
	CreatedTime       string     `json:"CreatedTime,omitempty"`
	SortBy		string	`json:"SortBy,omitempty"`
	OrderBy		string	`json:"OrderBy,omitempty"`
	OwnerIDs	[]int	`json:"OwnerIDs,omitempty"`
}

// This function is a handler that creates a POST API. <TBD>
//
// **Business Logic**: <TBD>
//
// Returns data as found, with a variable JSON Structure
func (h *Handler) GetTicketGrid(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t Grid_API
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	
	utils.ResponseAbstract(getTicketGrid(t),w)
}

func getTicketGrid(T Grid_API) []uint8{

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetTicketGrid").GetString("uri")


	url := felicitybaseurl + felicityapiuri

	j, err := json.Marshal(T)

	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)
	return utils.MakeHTTPPostCall(url,b)


}
