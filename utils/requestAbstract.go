package utils

import(
	"net/url"
	"net/http"
	"github.com/antigloss/go/logger"
	"encoding/json"
)

// Function that Abstracts out the Request to be executed for a POST Request
//
// Business Logic: Reads the Body coming in through an HTTP Request, Decodes the JSON into a Struct.
//
// Returns a Hash-Map of Variable Interface
func RequestAbstract(r *http.Request) map[string] interface{} {

        var t interface{}
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&t)
        if err != nil{
                logger.Error("Error in Decoding Request Body")
                logger.Error(err.Error())
        }
        defer r.Body.Close()
        return t.(map[string] interface{})
}


// Returns a hash-map of the GET Parameters of HTTP Request
func RequestAbstractGet(r *http.Request) url.Values {
	return r.URL.Query()

}
