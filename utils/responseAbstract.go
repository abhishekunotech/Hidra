package utils


import(
	"net/http"
	"encoding/json"
	"github.com/antigloss/go/logger"
)

// Function that sends the JSON response as a part of the HTTP Response
//
// Business Logic: Accepts a writer to the HTTP Response and a body as a stream of bytes, it will encode the stream as a JSON Object and sends up the JSON Object as a HTTP Response
//
// Returns Nothing
func ResponseAbstract(bodyText []uint8,w http.ResponseWriter){

	var data interface{}
        err := json.Unmarshal(bodyText, &data)

	

        if err != nil {
                logger.Error(err.Error())
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
}

