package handlers

import(
       "encoding/json" 
        "github.com/Unotechsoftware/Hydra/utils"
        "net/http"
)


// This function is a handler that generates a list of API
//
// **Business Logic**: Function reads the list of API from the configuration and gives out a list of API that are exposed by Hydra
//
// Returns a JSON array with Name,URI and Method of API
func (h *Handler) GetListOfAPIs(w http.ResponseWriter, r *http.Request) {
	listofapis(w)

}


func listofapis(w http.ResponseWriter){

	data := utils.ListRoutes()
	w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)




}

