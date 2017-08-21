package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func creatorOfTickets(jsonInput string) string {
	url := "http://192.168.2.108/felicity/nph-genericinterface.pl/Webservice/TicketAPI/TicketCreate?UserLogin=abhik&Password=abhik"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(jsonInput)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return string(body)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello bla")
}

func GetTicketDetails(w http.ResponseWriter, r *http.Request) {
	tick := &TicketDetails{TicketId: 123}
	json.NewEncoder(w).Encode(tick)
}

func GetCILogs(w http.ResponseWriter, r *http.Request) {

	url := "http://192.168.2.52:59200/_search?q=172.34.144.133&pretty=true&size=1"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Results: %v\n", data)

	json.NewEncoder(w).Encode(data)

}

func GetCIJobs(w http.ResponseWriter, r *http.Request) {
	CIJob := &CIJobs{JobId: 123}
	json.NewEncoder(w).Encode(CIJob)
}

func GetCIDetails(w http.ResponseWriter, r *http.Request) {
	CIDetail := &CIDetails{CIId: 123}
	json.NewEncoder(w).Encode(CIDetail)
}

func Ticketcreate(w http.ResponseWriter, r *http.Request) {

	bodyVal, _ := ioutil.ReadAll(r.Body)
	bodyValStrg := string(bodyVal)

	//fmt.Println("str is :",bodyValStrg)

	var jsonReturnString = creatorOfTickets(bodyValStrg)
	var jsonReturn = []byte(jsonReturnString)
	jsonRetVal, _ := json.Marshal(jsonReturn)
	var byteArr []byte
	base64.StdEncoding.Decode(byteArr, jsonRetVal)
	fmt.Fprintf(w, string(byteArr))
	//json.NewEncoder(w).Encode(Tick)
}
