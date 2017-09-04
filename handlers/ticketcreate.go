package handlers

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/base64"
	"encoding/json"
	"bytes"
)


type TicketResponseBody struct{
	

}

//Function to create ticket.
func Ticketcreate(w http.ResponseWriter, r *http.Request) {

        //ReadAll reads from response until an error or EOF and returns the data it read.
        bodyVal, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error Occured with Reading Body")
		fmt.Println(err.Error())
	}

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


func creatorOfTickets(jsonInput string) string {

        //API response is returned in JSON format from url
        url := "http://192.168.2.90:8080/felicity/nph-genericinterface.pl/Webservice/TicketAPI/TicketCreat?UserLogin=abhik&Password=abhik"
        fmt.Println("URL:>", url)
        //JSON input array
        var jsonStr = []byte(jsonInput)
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

        //Create custom header
        req.Header.Set("X-Custom-Header", "myvalue")
        req.Header.Set("Content-Type", "application/json")
        //A "Client" is an HTTP client.
        client := &http.Client{}

        //"Do" sends an HTTP request and returns an HTTP response.
        resp, err := client.Do(req)

        //Panic is a built-in function that stops the ordinary flow of control and begins panicking.
        //Panics can be initiated by invoking panic directly. They can also be caused by runtime errors.
        //Errors are handled if any.
        if err != nil {
                panic(err)
        }
        //When "err" is nil, "resp" always contains a non-nil "resp.Body".
        //Callers should close "resp.Body" using defer when done reading from it.
        defer resp.Body.Close()

        fmt.Println("response Status:", resp.Status)
        fmt.Println("response Headers:", resp.Header)
        //ReadAll reads from response until an error or EOF and returns the data it read.

        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println("response Body:", string(body))
        return string(body)
}
