package handlers

import(


	"net/http"
	"io/ioutil"
	"encoding/json"
	"io"
	"github.com/antigloss/go/logger"
)

type SessionObject struct{
        SessionIDStrg   string  `json:"SessionID"`
}

func callSessionDetails(username string, password string) string{
        url := "http://192.168.2.90:8080/felicity/nph-genericinterface.pl/Webservice/SessionAPI/SessionCreate?UserLogin="+username+"&Password="+password
       
        client := &http.Client{}
        var bodyReader io.Reader
        req, err := http.NewRequest("GET", url, bodyReader)
        resp, err := client.Do(req)
        if err != nil{
                logger.Error("\n\n Session Creation failed because - \n\n")
                logger.Error(err.Error())
        }
        req.Close = true
        bodyText, err := ioutil.ReadAll(resp.Body)
        var data SessionObject
        err = json.Unmarshal(bodyText, &data)
        if err != nil{
                logger.Error("\n\n Error Occured in unmarshalling Session JSON \n\n")
                logger.Error(err.Error())
        }
        return data.SessionIDStrg
}

