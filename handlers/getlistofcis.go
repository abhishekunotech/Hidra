package handlers

import (
	"encoding/json"
	"fmt"
//	"github.com/antigloss/go/logger"
//	"io"
	"io/ioutil"
	"net/http"
	"github.com/Unotechsoftware/Hydra/lerna"
)
type ListOfCIs struct {

        ZeroCI []ZeroCIArray
      OneCI []OneCIArray 
}

type ZeroCIArray struct{
	DeplState	string	`json:"DeplState"`
	CI_Id	 	int	`json:"ci_id"`
	Create_Time 	string	`json:"create_time"`
	Link_Name  	string  `json:"link_name"`
        ClassName   	string  `json:"ClassName"`
        TypeId       	int     `json:"type_id"`
	ClassId 	int     `json:"class_id"`
        InciState       string  `json:"InciState"`
        Configitem_Number string  `json:"configitem_number"`
	TicketId	int        `json:"ticket_id"`
        Name    string  `json:"Name"`
	
}
type OneCIArray struct{
	DeplState       string  `json:"DeplState"`
        CI_Id           int     `json:"ci_id"`
        Create_Time     string  `json:"create_time"`
        Link_Name       string  `json:"link_name"`
        ClassName       string  `json:"ClassName"`
        TypeId          int     `json:"type_id"`
        ClassId         int     `json:"class_id"`
        InciState       string  `json:"InciState"`
        Configitem_Number string  `json:"configitem_number"`
        TicketId        int        `json:"ticket_id"`
        Name    string  `json:"Name"`
}

func GetListOfCIs(w http.ResponseWriter, r *http.Request) {

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.getticketinfo").GetString("uri")
	felicityusername := ConfObj.Sub("components.otrs.apis.getticketinfo.parameters").GetString("UserLogin")
	felicitypassword := ConfObj.Sub("components.otrs.apis.getticketinfo.parameters").GetString("Password")
        url := felicitybaseurl+felicityapiuri+"?UserLogin="+felicityusername+"&Password="+felicitypassword
	fmt.Println(url)


	//API response is returned in JSON format from url

	//url := "http://192.168.2.52:59200/_search?q=172.34.144.133&pretty=true&size=1"

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
