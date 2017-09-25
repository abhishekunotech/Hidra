package utils


import(
	"net/http"
	"encoding/json"
	"github.com/antigloss/go/logger"
)


func ResponseAbstract(bodyText []uint8,w http.ResponseWriter){

	 var data interface{}
        err := json.Unmarshal(bodyText, &data)
        if err != nil {
                logger.Error(err.Error())
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
}

