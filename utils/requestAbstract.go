package utils

import(
	"net/http"
	"github.com/antigloss/go/logger"
	"encoding/json"
	"net/url"
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


func RequestAbstractGet(r *http.Request) url.Values {
	return r.URL.Query()

}

// Returns a hash-map of the GET Parameters of HTTP Request
func RequestAbstractGet1(r *http.Request) string {
	//return r.URL.Query()
	var mapHttp url.Values
	mapHttp = RequestAbstractGet(r)
        actionStrg := ""
        count := 0
        for key,value := range mapHttp{
        if count == 0{
                actionStrg = actionStrg + "?" + key +"="+value[0]

        }else{
                actionStrg = actionStrg + "&" + key + "="+ value[0]
        }
        count = count + 1
        }

	return actionStrg
}
