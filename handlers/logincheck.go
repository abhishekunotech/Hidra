package handlers

import (
	"fmt"
	"encoding/json"
	//"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/Unotechsoftware/pearly-gates-go"
	"github.com/antigloss/go/logger"
//	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
)

type User_Req struct {
	Username	string	`json:"Username"`
	Password	string	`json:"Password"`
}

func (h *Handler) LoginCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In logincheck")
	decoder := json.NewDecoder(r.Body)
	var t User_Req
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	fmt.Println(t.Username)
	fmt.Println(t.Password)

	w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(t)

}
