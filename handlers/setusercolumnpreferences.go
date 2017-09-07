package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/antigloss/go/logger"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
	"bytes"
	"io/ioutil"
)

type SUCP_Request struct{
	UserLogin	string	`json:"UserLogin"`
	Password	string	`json:"Password"`
	Action		string	`json:"Action"`
	ColumnSelected	[]string	`json:"ColumnSelected"`
}


func SetUserColumnPreferences(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t SUCP_Request
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println("Error Occured")
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
		fmt.Println(err.Error())
	}
	defer r.Body.Close()
	
	setUserColumnPreferences(t,w,r)
}

func setUserColumnPreferences(T SUCP_Request, w http.ResponseWriter, r *http.Request) {

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

        felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
        felicityapiuri :=  ConfObj.Sub("components.otrs.apis.setusercolumnpreferences").GetString("uri")
        //sessionIDString := callSessionDetails(T.UserLogin,T.Password)
	
	//fmt.Println("SESSION ID IS")
	//fmt.Println(sessionIDString) 	
	
	//url := felicitybaseurl+felicityapiuri+"?SessionID="+sessionIDString
	url := felicitybaseurl+felicityapiuri+"?UserLogin="+T.UserLogin+"&Password="+T.Password
	j, err := json.Marshal(T)
	
	if err!=nil {
		fmt.Println("Error in Marshaling JsON")
		fmt.Println(err.Error())
	}

	b := bytes.NewBuffer(j)

	fmt.Println(b)
/*
        resp, err := http.Post(url, "application/json", b)

        if err != nil {
                fmt.Println("Error OCcured")
                fmt.Println(err.Error())
        }

        defer resp.Body.Close()

        fmt.Println("response Status:", resp.Status)
        fmt.Println("response Headers:", resp.Header)

        var bodyText []byte
        var data interface{}
        err = json.Unmarshal(bodyText, &data)
        if err != nil {
                fmt.Println("JSON  Unmarshalling failed")
		fmt.Println(err.Error())
		logger.Error(err.Error())
        }
	fmt.Println(bodyText)
*/

	 client := &http.Client{}
        
        req, err := http.NewRequest("POST", url, b)
	
	if err != nil {
		fmt.Println("\n\n Request to Create Request Failed \n\n")
		fmt.Println(err.Error())
	}

	fmt.Println("Request")

	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	fmt.Println(req)
        resp, err := client.Do(req)
        if err != nil{
                fmt.Println("\n\n POST REQUEST TO FELICITY FAILED \n\n")
               	fmt.Println(err.Error())
        }
        //req.Close = true
        bodyText, err := ioutil.ReadAll(resp.Body)
        var data interface{}
        err = json.Unmarshal(bodyText, &data)
        if err != nil{
                logger.Error("\n\n Error Occured in unmarshalling Session JSON \n\n")
                logger.Error(err.Error())
        }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}
