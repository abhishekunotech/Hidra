package handlers

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/antigloss/go/logger"
	"bytes"
	"time"
	"fmt"
)

/*
{"ArticleID":"2927","TicketNumber":"2017090410000024","TicketID":"629"}

*/
type TicketResponseBody struct{
	ArticleID string	`json:"ArticleID,omitempty"`
	TicketNumber	string	`json:"TicketNumber"`
	TicketID	string	`json:"TicketID"`
}

//Function to create ticket.
func (h *Handler)Ticketcreate(w http.ResponseWriter, r *http.Request) {

        //ReadAll reads from response until an error or EOF and returns the data it read.
        bodyVal, err := ioutil.ReadAll(r.Body)

	if err != nil {
		logger.Error("Error Occured with Reading Body")
		logger.Error(err.Error())
	}

        bodyValStrg := string(bodyVal)

        //Function call to create ticket and get the response
        creatorOfTickets(bodyValStrg,w,r)
        
	/*var jsonReturn = []byte(jsonReturnString)
        //Decode JSON response
        jsonRetVal, _ := json.Marshal(jsonReturn)
        var byteArr []byte
        //Display the response
        base64.StdEncoding.Decode(byteArr, jsonRetVal)
        fmt.Fprintf(w, string(byteArr))
        //json.NewEncoder(w).Encode(Tick)*/
}


func creatorOfTickets(jsonInput string,w http.ResponseWriter,r *http.Request) {


	 http.DefaultClient.Timeout = 10 * time.Second
	ConfObj := lerna.ReturnConfigObject()
	
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
        felicityapiuri :=  ConfObj.Sub("components.otrs.apis.ticketcreate").GetString("uri")
        passwordStrg := ConfObj.Sub("components.otrs.apis.ticketcreate.parameters").GetString("Password")	
	usernameStrg := ConfObj.Sub("components.otrs.apis.ticketcreate.parameters").GetString("Username")
	sessionIDString := callSessionDetails(usernameStrg,passwordStrg)	
	

        url := felicitybaseurl+felicityapiuri+"?SessionID="+sessionIDString
        fmt.Println("URLi:>", url)

	jsonStr1 := bytes.NewBufferString(jsonInput)
	
	resp, err := http.Post(url, "application/json", jsonStr1)
	
        if err != nil {
		fmt.Println("Error OCcured")
		fmt.Println(err.Error())
        }
        defer resp.Body.Close()

        fmt.Println("response Status:", resp.Status)
        fmt.Println("response Headers:", resp.Header)

        body, _ := ioutil.ReadAll(resp.Body)
	var bodyText []byte
	var data TicketResponseBody
        err = json.Unmarshal(bodyText, &data)
        if err != nil {
                logger.Error(err.Error())
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)	

        fmt.Println("response Body:", string(body))
}
