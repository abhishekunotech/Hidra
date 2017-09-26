package handlers

import(
	"fmt"
       "encoding/json" 
        "github.com/Unotechsoftware/Hydra/utils"
        "net/http"
)

func (h *Handler) GetListOfAPIs(w http.ResponseWriter, r *http.Request) {
	listofapis(w)

}


func listofapis(w http.ResponseWriter){

	data := utils.ListRoutes()
	fmt.Println(data)
	w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)




}

