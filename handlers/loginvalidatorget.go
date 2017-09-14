package handlers

import (
	"encoding/json"
	"net/http"
)

type UsernamefromJWT struct{
	Username	string	`json:"Username"`
}


func loginValidatorGet(w http.ResponseWriter, r *http.Request, token string) {

	//Read Token
	//Decode Token
	//GET USERNAME
	//Store UserName into Return Struct "data"

	var data UsernamefromJWT
	data.Username = token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

//Function to get list of work orders
// Request as http://ip-host/getListOfWorkOrders?ticketID=521&password=abhik&userLogin=abhik

func (h *Handler) LoginValidatorGet(w http.ResponseWriter, r *http.Request) {
	mapHttp := r.URL.Query()
	var token string
	for key, value := range mapHttp {
		if key == "Token" {
			for _, valueStrg := range value {
				token = valueStrg
			}
		}
	}

	loginValidatorGet(w,r,token)
}
