package handlers

import (
	"encoding/json"
	"github.com/antigloss/go/logger"
	"net/http"
)

type User_Req struct {
	Username	string	`json:"Username"`
	Password	string	`json:"Password"`
}
// LoginCheck provides basic user authentication functionality.
//
// **Business Logic**: Function takes Username and Password of the user as an input in JSON Body
func (h *Handler) LoginCheck(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t User_Req
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()


	w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(t)

}
