package handlers

import (
	"encoding/json"
	"github.com/antigloss/go/logger"
	"net/http"
)

type User_Request struct {
	Username	string	`json:"Username"`
	Password	string	`json:"Password"`
}

func (h *Handler) IsValidUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t User_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()


	w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(t)

}
