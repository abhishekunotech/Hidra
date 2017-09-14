package handlers

import (
	"encoding/json"
	//"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	//"io/ioutil"
	"github.com/Unotechsoftware/Hydra/utils"
	"net/http"
	"gopkg.in/dgrijalva/jwt-go.v2"
)

type Token_Request struct {
	Username      string   `json:"Username"`
	Password       string   `json:"Password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t Token_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	tokenGenerator(t, w, r)
}

func tokenGenerator(T Token_Request, w http.ResponseWriter, r *http.Request) {


	someBytes := utils.ReadFile("/root/.ssh/private.pem")
	token := jwt.New(jwt.SigningMethodRS256)  
	token.Claims["Username"] = T.Username
	token.Claims["Password"] = T.Password
	
	signKey,_ := jwt.ParseRSAPrivateKeyFromPEM(someBytes)

	tokenString, _ := token.SignedString(signKey)

	
	//Create a variable 'data' that holds the JWT
	
	var data JWToken
	
	data.Token = tokenString

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}
