/*
    This is the List of Handler Functions

    This package provides the source code for all the APIs exposed by Hydra.

*/
// Author: Operations Management Team - Unotech Software.
package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
}

//Structure for go type definition of the JSON
type TicketDetails struct {
	TicketId int
}

type CIDetails struct {
	CIId int
}

type CILogs struct {
	LogId int
}

type CIJobs struct {
	JobId int
}

type Logs struct {
	Took int `json:"took"`

	Timed_out bool `json:"timed_out"`

	Shards []Shards_info

	Hits []Hits_info
}

type Shards_info struct {
	Total int `json:"total"`

	Successful int `json:"successful"`

	Failed int `json:"failed"`
}

type Hits_info struct {
	Total int `json:"total"`

	Max_score int `json:"max_score"`

	Hits []Hits_array
}

type Hits_array struct {
	Index string `json:"_index"`

	Type string `json:"_type"`

	Score string `json:"_score"`

	Source []Source_array
}

type Source_array struct {
	Nagios_type string `json:"nagios_type"`

	Nagios_hostname string `json:"nagios_hostname"`

	Nagios_service string `json:"nagios_service"`

	Nagios_state string `json:"nagios_stat"`

	Nagios_statetype string `json:"nagios_statetype"`

	Nagios_statecode string `json:"nagios_statecode"`
}
type getCIDetails struct {
	TotalCount int `json:"totalcount"`

	Count int `json:"count"`

	Sort int `json:"sort"`

	Order string `json:"order"`

	Data []DataArray

	Content_Range string `json:"content-range"`
}

type DataArray struct {
	Zero []ZeroArray
	One  []OneArray
}

type ZeroArray struct {
	One         string `json:"1"`
	Three       string `json:"3"`
	Four        string `json:"4"`
	Five        string `json:"5"`
	Seventeen   string `json:"17"`
	Nineteen    string `json:"19"`
	TwentyThree string `json:"23"`
	ThirtyOne   string `json:"31"`
	Forty       string `json:"40"`
	FortyFive   string `json:"45"`
	OneTwoSix   string `json:"126"`
	IP          string `json:"10008"`
}
type OneArray struct {
	One         string `json:"1"`
	Three       string `json:"3"`
	Four        string `json:"4"`
	Five        string `json:"5"`
	Seventeen   string `json:"17"`
	Nineteen    string `json:"19"`
	TwentyThree string `json:"23"`
	ThirtyOne   string `json:"31"`
	Forty       string `json:"40"`
	FortyFive   string `json:"45"`
	Ots         []OtsArray
	IP          string `json:"10008"`
}

type OtsArray struct {
	Zero  string `json:"0"`
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
}

// Function to get CI logs.
func (h *Handler) GetCILogs(w http.ResponseWriter, r *http.Request) {

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.graylog").GetString("url")
	felicityapiuri := ConfObj.Sub("components.graylog.apis.getcilogs").GetString("uri")
	ipStrg := ConfObj.Sub("components.graylog.apis.getcilogs.parameters").GetString("ip")
	prettyStrg := ConfObj.Sub("components.graylog.apis.getcilogs.parameters").GetString("pretty")

	sizeStrg := ConfObj.Sub("components.graylog.apis.getcilogs.parameters").GetString("size")
	url1 := felicitybaseurl + felicityapiuri + ipStrg + "&pretty=" + prettyStrg + "&size=" + sizeStrg


	body := utils.MakeHTTPGetCall(url1)
	var data interface{}

	//To decode JSON data, use the Unmarshal function.
	err := json.Unmarshal(body, &data)
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Printf("Results: %v\n", data)
	//Encode the data
	json.NewEncoder(w).Encode(data)

}

//Function to get CI jobs.
func (h *Handler) GetCIJobs(w http.ResponseWriter, r *http.Request) {
	CIJob := &CIJobs{JobId: 123}
	json.NewEncoder(w).Encode(CIJob)
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.rundeck").GetString("url")
	felicityapiuri := ConfObj.Sub("components.rundeck.apis.getcijobs").GetString("uri")
	felicityusername := ConfObj.Sub("components.rundeck.apis.getcijobs.parameters").GetString("UserLogin")
	felicitypassword := ConfObj.Sub("components.rundeck.apis.getcijobs.parameters").GetString("Password")
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + felicityusername + "&Password=" + felicitypassword
	fmt.Println(url)
}

//Function to get CI details.
func (h *Handler) GetCIDetails(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	mapHttp := r.URL.Query()
	var userName string
	var password string
	var hostip string
	for key, value := range mapHttp {

		if key == "login" {
			for _, valueStrg := range value {
				userName = valueStrg
			}
		}

		if key == "password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}

		if key == "host_ip" {
			for _, valueStrg := range value {
				hostip = valueStrg
			}
		}
	}

	s := basicAuth(userName, password)

	var f interface{}
	session_t := json.Unmarshal([]byte(s), &f)
	if session_t != nil {
		fmt.Println("Error occured ", session_t)
	}
	m := f.(map[string]interface{})
	var session_string string
	for k, v := range m {
		if k == "session_token" {
			session_string = "" + v.(string)
		}
	}

	url := "http://192.168.2.72/glpi/apirest.php/search/Computer?criteria[0][link]=AND&criteria[0][itemtype]=Computer&criteria[0][field]=10008&criteria[0][searchtype]=contains&criteria[0][value]=" + hostip + "&session_token=" + session_string

	fmt.Println("URL is " + url)
	res, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err.Error())
	}
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Printf("Results: %v\n", data)
	json.NewEncoder(w).Encode(data)
}

//Function to set authentication header for CIDetails.
func basicAuth(username string, password string) string {

	client := &http.Client{}
	encoded_login := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	fmt.Printf(username)
	fmt.Printf(password)
	fmt.Printf(encoded_login)
	req, err := http.NewRequest("GET", "http://192.168.2.72/glpi/apirest.php/initSession", nil)
	//req.SetBasicAuth(username,password)
	req.Header.Set("Authorization", "Basic Z2xwaTpnbHBp")
	resp, err := client.Do(req)
	log.Println("Meow")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Meow ends")
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	fmt.Printf(s)
	return s
}
