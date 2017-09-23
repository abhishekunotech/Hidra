package handlers

import (
	"fmt"
	"encoding/json"
//	"strconv"
	//"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
//	//"io/ioutil"
//	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
//	"gopkg.in/dgrijalva/jwt-go.v2"
	//"github.com/codegangsta/negroni"
	//"github.com/dgrijalva/jwt-go"	
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

	fmt.Println(t.Username)
	fmt.Println(t.Password)

	w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(t)

}
