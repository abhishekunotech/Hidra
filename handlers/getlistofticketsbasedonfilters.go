package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"net/http"
)

type Grid_API struct {
	UserLogin         string   `json:"UserLogin"`
	Password          string   `json:"Password"`
	PageSize          int      `json:"PageSize"`
	PageNumber        int      `json:"pageNumber"`
	Action            string   `json:"Action"`
	TicketNumber      string   `json:"TicketNumber"`
	TypeID            int      `json:"TypeID"`
	QueueIDs          []int    `json:"QueueIDs"`
	StateIDs          []int    `json:"StateIDs"`
	CustomerUserLogin []string `json:"CustomerUserLogin"`
	PriorityIDs       []int    `json:"PriorityIDs"`
	CreatedUserIDs    []int    `json:"CreatedUserIDs"`
	SourceIDs         []int    `json:"SourceIDs"`
	Due               string   `json:"Due"`
	GroupIDs          []int    `json:"GroupIDs"`
	CreatedTime       int      `json:"CreatedTime"`
}

func (h *Handler) GetTicketGrid(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t Grid_API
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println("Error Occured")
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
		fmt.Println(err.Error())
	}
	defer r.Body.Close()

	GetTicketGrid(t, w, r)
}

func GetTicketGrid(T Grid_API, w http.ResponseWriter, r *http.Request) {

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetTicketGrid").GetString("uri")
	//sessionIDString := callSessionDetails(T.UserLogin,T.Password)

	//fmt.Println("SESSION ID IS")
	//fmt.Println(sessionIDString)

	//url := felicitybaseurl+felicityapiuri+"?SessionID="+sessionIDString
	//url := felicitybaseurl+felicityapiuri+"?UserLogin="+T.UserLogin+"&Password="+T.Password
	url := felicitybaseurl + felicityapiuri

	fmt.Println("\n\n\n URL IS ")
	fmt.Println(url)
	j, err := json.Marshal(T)

	if err != nil {
		fmt.Println("Error in Marshaling JsON")
		fmt.Println(err.Error())
	}

	b := bytes.NewBuffer(j)

	fmt.Println(b)
	/*
	           resp, err := http.Post(url, "application/json", b)
	           if err != nil {
	                   fmt.Println("Error OCcured here")
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
	if err != nil {
		fmt.Println("\n\n POST REQUEST TO FELICITY FAILED \n\n")
		fmt.Println(err.Error())
	}
	//req.Close = true
	bodyText, err := ioutil.ReadAll(resp.Body)
	var data interface{}
	err = json.Unmarshal(bodyText, &data)
	if err != nil {
		logger.Error("\n\n Error Occured in unmarshalling Session JSON \n\n")
		logger.Error(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}
