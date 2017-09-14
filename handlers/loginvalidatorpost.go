package handlers

import (
	"fmt"
	//"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/antigloss/go/logger"
	//"io/ioutil"
	"net/http"
	 "gopkg.in/dgrijalva/jwt-go.v2"
)

type JWToken struct {
	Token      string   `json:"Token"`
}

func (h *Handler) LoginValidatorPost(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t JWToken
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	decodeJWTPost(t, w, r)
}

func decodeJWTPost(T JWToken, w http.ResponseWriter, r *http.Request) {


	fmt.Println("Token is ")
	fmt.Println(T.Token)
	//Read the JWToken

	token, err := jwt.Parse(T.Token,func(token *jwt.Token) (interface{}, error) {
	verifyBytes := utils.ReadFile("/root/.ssh/public.pem")	
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		fmt.Println("I dont necessarily understand GO")
		fmt.Println(err.Error())
	}
		return verifyKey, nil
	})

	if err != nil {
		fmt.Println("SOme Error occured in Parsing")
		fmt.Println("I love go, because it lets me nest my errors")
		fmt.Println(err.Error())
	}
	//fmt.Println(token.Claims)
	
	usernameVal := token.Claims["Username"].(string)

	var data UsernamefromJWT

	data.Username = usernameVal

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}
