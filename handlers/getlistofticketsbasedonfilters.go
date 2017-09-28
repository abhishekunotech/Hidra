package handlers



import (
	"github.com/Unotechsoftware/Hydra/utils"
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)

type Grid_API struct {
	UserLogin         string   `json:"UserLogin"`
	Password          string   `json:"Password"`
	PageSize          int      `json:"PageSize"`
	PageNumber        int      `json:"pageNumber"`
	Action            string   `json:"Action"`
	TicketNumber      string   `json:"TicketNumber"`
	TypeID            int      `json:"TypeID"`
	QueueIDs          []int    `json:"QueueIDs"`
	StateIDs          []int    `json:"StateIDs"`
	CustomerUserLogin []string `json:"CustomerUserLogin"`
	PriorityIDs       []int    `json:"PriorityIDs"`
	CreatedUserIDs    []int    `json:"CreatedUserIDs"`
	SourceIDs         []int    `json:"SourceIDs"`
	Due               string   `json:"Due"`
	GroupIDs          []string    `json:"GroupIDs"`
	CreatedTime       string     `json:"CreatedTime"`
	SortBy		string	`json:"SortBy"`
	OrderBy		string	`json:"OrderBy"`
}

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
