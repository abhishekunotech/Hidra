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
	"log"
	//"net"
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
 type getCIDetails struct {
	TotalCount 	int `json:"totalcount"`

	Count int `json:"count"`

	Sort int `json:"sort"`

	Order string `json:"order"`
 	
	Data []DataArray	

	Content_Range string `json:"content-range"`

}

type DataArray struct {

        Zero[]ZeroArray
      One []OneArray 
}

type ZeroArray struct {

	One string `json:"1"`
	Three string `json:"3"`
	Four string `json:"4"`
	Five string `json:"5"`
	Seventeen string `json:"17"`
	Nineteen string `json:"19"`
	TwentyThree string `json:"23"`	
	 ThirtyOne string `json:"31"`
	 Forty string `json:"40"`
	 FortyFive string `json:"45"`
	 OneTwoSix string `json:"126"`
	 IP  string `json:"10008"`
}
 type OneArray struct {
        One string `json:"1"`
        Three string `json:"3"`
        Four string `json:"4"`
        Five string `json:"5"`
        Seventeen string `json:"17"`
        Nineteen string `json:"19"`
        TwentyThree string `json:"23"`
         ThirtyOne string `json:"31"`
         Forty string `json:"40"`
         FortyFive string `json:"45"`
         Ots []OtsArray
         IP  string `json:"10008"` 
          
}       

type OtsArray struct {

        Zero string `json:"0"`
        One string `json:"1"`
        Two string `json:"2"`
        Three string `json:"3"`
        Four string `json:"4"`
}
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

   body, _ := ioutil.ReadAll(r.Body)
        mapHttp := r.URL.Query()
        var userName string
        var password string
        var hostip string
        for key,value := range mapHttp {
             
                if key == "login"{
                        for _, valueStrg := range value {
                                userName = valueStrg
                        }
                }

                if key == "password"{
                        for _, valueStrg := range value {
                                password = valueStrg
                        }
                }

		if key == "host_ip"{
			for _, valueStrg := range value {
				hostip = valueStrg	
			}
		}
        }

	fmt.Println("Username is "+userName)
	fmt.Println("Password is "+password)
	/*
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Errorf("ip: %q is not IP:port", r.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
	}
	*/

	/*body, _ := ioutil.ReadAll(r.Body)

        mapHttp := r.URL.Query()

        for key,value := range mapHttp {
                fmt.Fprintf(w,key+" is key \n")
                for _, valueStrg := range value {
                        fmt.Fprintf(w,valueStrg)
                        fmt.Fprintf(w," is value \n")
                }
 }
        bodyStrg := string(body[:])

        fmt.Fprintf(w,"www"+bodyStrg+"\n")
*/
     	s := basicAuth(userName , password)

var f interface{}
session_t := json.Unmarshal([]byte(s), &f)
if session_t != nil {
       fmt.Println("Error occured ",session_t)
}
m := f.(map[string]interface{})
var session_string string 
       for k,v := range m{
       if k == "session_token"{
               session_string = ""+v.(string)
       }
}

url := "http://192.168.2.72/glpi/apirest.php/search/Computer?criteria[0][link]=AND&criteria[0][itemtype]=Computer&criteria[0][field]=10008&criteria[0][searchtype]=contains&criteria[0][value]="+hostip+"&session_token="+session_string

fmt.Println("URL is "+url)
res, err := http.Get(url)
    if err != nil {
        panic(err.Error())
    }
    
body, err = ioutil.ReadAll(res.Body)
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

//Function to set authentication header for CIDetails.
func basicAuth(username string, password string) string {
    
    client := &http.Client{}
    encoded_login := base64.StdEncoding.EncodeToString([]byte(username+":"+password))
    fmt.Printf(username)
    fmt.Printf(password)
    fmt.Printf(encoded_login)
    req, err := http.NewRequest("GET", "http://192.168.2.72/glpi/apirest.php/initSession",nil)
    //req.SetBasicAuth(username,password)
    req.Header.Set("Authorization", "Basic Z2xwaTpnbHBp")
    resp, err := client.Do(req)
    log.Println("Meow")
    if err != nil{
        log.Fatal(err)
    }
    log.Println("Meow ends")
    bodyText, err := ioutil.ReadAll(resp.Body)
    s := string(bodyText)
fmt.Printf(s)
    return s
}
