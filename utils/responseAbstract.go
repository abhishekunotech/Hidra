package utils


import(
	"net/http"
	"encoding/json"
)


func ResponseAbstract(t interface{},w http.ResponseWriter){
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(t)
}

