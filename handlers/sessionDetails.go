package handlers

import(
	"fmt"
	"reflect"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"io"
)

type SessionObject struct{
        SessionIDStrg   string  `json:"SessionID"`
}

func callSessionDetails(username string, password string) string{
        url := "http://192.168.2.90:8080/felicity/nph-genericinterface.pl/Webservice/SessionAPI/SessionCreate?UserLogin="+username+"&Password="+password
        fmt.Printf("\n\n url is "+url+"\n\n")
        client := &http.Client{}
        var bodyReader io.Reader
        req, err := http.NewRequest("GET", url, bodyReader)
        resp, err := client.Do(req)
        if err != nil{
                fmt.Printf("\n\n Session Creation failed because - \n\n")
                fmt.Printf(err.Error())
        }
        req.Close = true
        bodyText, err := ioutil.ReadAll(resp.Body)
        var data SessionObject
        err = json.Unmarshal(bodyText, &data)
        fmt.Printf("\n\n body is ")
        fmt.Printf(reflect.TypeOf(data).String())
        fmt.Printf("\n\n")
        if err != nil{
                fmt.Printf("\n\n Error Occured in unmarshalling Session JSON \n\n")
                fmt.Printf(err.Error())
        }
        fmt.Printf("\n\nSESSION ID IN CALL SESSION DETAILS IS "+data.SessionIDStrg+"\n\n")
        return data.SessionIDStrg
}

