package utils

import(
	"net/url"
	"net/http"
	"github.com/antigloss/go/logger"
	"encoding/json"
)


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
