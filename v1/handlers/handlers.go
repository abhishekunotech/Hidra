//Package handlers contains definitions for different functions.
//Install using go install in this directory.
//Author: Operations Management Team - Unotech Software.

package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

type Ticket struct {
	ArticleID    string
	TicketNumber string
	TicketID     string
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

//This function will generate tickets and return the ticket in json format.
func creatorOfTickets(jsonInput string) string {

//API response is returned in JSON format from url 
	url := "http://192.168.2.108/felicity/nph-genericinterface.pl/Webservice/TicketAPI/TicketCreate?UserLogin=abhik&Password=abhik"
	fmt.Println("URL:>", url)
//JSON input array
	var jsonStr = []byte(jsonInput)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

//Create custom header 
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
//A “Client” is an HTTP client.
	client := &http.Client{}

//“Do” sends an HTTP request and returns an HTTP response.
	resp, err := client.Do(req)

//Panic is a built-in function that stops the ordinary flow of control and begins panicking.
//Panics can be initiated by invoking panic directly. They can also be caused by runtime errors.
//Errors are handled if any.
	if err != nil {
		panic(err)
	}
//When “err” is nil, “resp” always contains a non-nil “resp.Body”. 
//Callers should close “resp.Body” using defer when done reading from it.
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
//ReadAll reads from response until an error or EOF and returns the data it read.

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return string(body)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello bla")
}
//Function to get the details about ticket.
func GetTicketDetails(w http.ResponseWriter, r *http.Request) {
	tick := &TicketDetails{TicketId: 123}
	json.NewEncoder(w).Encode(tick)
}
//Function to get CI logs.
func GetCILogs(w http.ResponseWriter, r *http.Request) {

//API response is returned in JSON format from url

	url := "http://192.168.2.52:59200/_search?q=172.34.144.133&pretty=true&size=1"
	res, err := http.Get(url)

//Errors are handled if any. 
	if err != nil {
		panic(err.Error())
	}

//ReadAll reads from response until an error or EOF and returns the data it read.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var data interface{}

//To decode JSON data, use the Unmarshal function. 
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Results: %v\n", data)
//Encode the data
	json.NewEncoder(w).Encode(data)

}
//Function to get CI jobs.
func GetCIJobs(w http.ResponseWriter, r *http.Request) {
	CIJob := &CIJobs{JobId: 123}
	json.NewEncoder(w).Encode(CIJob)
}

//Function to get CI details.
func GetCIDetails(w http.ResponseWriter, r *http.Request) {
	CIDetail := &CIDetails{CIId: 123}
	json.NewEncoder(w).Encode(CIDetail)
}
//Function to create ticket.
func Ticketcreate(w http.ResponseWriter, r *http.Request) {

//ReadAll reads from response until an error or EOF and returns the data it read.
	bodyVal, _ := ioutil.ReadAll(r.Body)
	bodyValStrg := string(bodyVal)

//Function call to create ticket and get the response
	var jsonReturnString = creatorOfTickets(bodyValStrg)
	var jsonReturn = []byte(jsonReturnString)
//Decode JSON response
	jsonRetVal, _ := json.Marshal(jsonReturn)
	var byteArr []byte
//Display the response
	base64.StdEncoding.Decode(byteArr, jsonRetVal)
	fmt.Fprintf(w, string(byteArr))
	//json.NewEncoder(w).Encode(Tick)
}
