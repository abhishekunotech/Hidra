package handlers

import(
       "encoding/json" 
        "github.com/Unotechsoftware/Hydra/utils"
        "net/http"
)

func (h *Handler) GetListOfAPIs(w http.ResponseWriter, r *http.Request) {
	listofapis(w)

}


func listofapis(w http.ResponseWriter){

	data := utils.ListRoutes()
	w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)




}

