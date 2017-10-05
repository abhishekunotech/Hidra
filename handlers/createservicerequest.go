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
	UserLogin      string   `json:"UserLogin"`
	Password       string   `json:"Password"`
	Ticket         Ticket  `json:"Ticket"`
	Article        Article  `json:"Article"`
}
type Ticket struct{
	Title		 string   `json:"Title"`
	State      	 string   `json:"State"`
	Priority         string   `json:"Priority"`
	Queue  	 	 string   `json:"Queue"`
	CustomerUser     string   `json:"CustomerUser"`
	Type      	 string   `json:"Type"`
	ServiceID        string   `json:"ServiceID"`
	SLAID            string   `json:"SLAID"`
	Owner            string   `json:"Owner"`
	Responsible      string   `json:"Responsible"`
}
type Article struct{
	Subject      string   `json:"Subject"`
	Body         string   `json:"Body"`
	Charset      string   `json:"Charset"`
	MimeType     string   `json:"MimeType"`
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
