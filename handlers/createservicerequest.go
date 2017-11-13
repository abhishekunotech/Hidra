package handlers

import (
	"fmt"
	"bytes"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"net/http"
)



// This Type defines the Input JSON Body for Creating a Service Request
//
// It includes sub-types Ticket and Article
type SR_Request struct {
	UserLogin      string   `json:"UserLogin,omitempty"`
	CustomerUserLogin	string	`json:"CustomerUserLogin,omitempty"`
	Password       string   `json:"Password,omitempty"`
	Ticket         *Ticket  `json:"Ticket,omitempty"`
	Article        *Article  `json:"Article,omitempty"`
}
type Ticket struct{
	Title		 string   `json:"Title,omitempty"`
	State      	 string   `json:"State,omitempty"`
	Priority         string   `json:"Priority,omitempty"`
	Queue  	 	 string   `json:"Queue,omitempty"`
	CustomerUser     string   `json:"CustomerUser,omitempty"`
	Type      	 string   `json:"Type,omitempty"`
	ServiceID        string   `json:"ServiceID,omitempty"`
	SLAID            string   `json:"SLAID,omitempty"`
	Owner            string   `json:"Owner,omitempty"`
	Responsible      string   `json:"Responsible,omitempty"`
}
type Article struct{
	Subject      string   `json:"Subject,omitempty"`
	Body         string   `json:"Body,omitempty"`
	Charset      string   `json:"Charset,omitempty"`
	MimeType     string   `json:"MimeType,omitempty"`
}


// This function is a handler that creates a Service Request based on Input from UI
//
// **Business Logic**: Function takes as an input a JSON Body and uses the Ticket and Article in Request Body to generate a Service Request
//
// Returns data as shown in examples
func (h *Handler) CreateServiceRequest(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t SR_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	createServiceRequest(t, w, r)
}

func createServiceRequest(T SR_Request, w http.ResponseWriter, r *http.Request) {

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.CreateServiceRequest").GetString("uri")
	fmt.Println(T)
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + T.UserLogin + "&Password=" + T.Password
	j, err := json.Marshal(T)
	fmt.Println(url)
	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}
	fmt.Println(j)
	b := bytes.NewBuffer(j)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, b)

	if err != nil {
		logger.Error("\n\n Request to Create Request Failed \n\n")
		logger.Error(err.Error())
	}

	logger.Info("Request")

	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("\n\n POST REQUEST TO FELICITY FAILED \n\n")
		logger.Error(err.Error())
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
