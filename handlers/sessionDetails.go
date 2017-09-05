package handlers

import(
	"github.com/Unotechsoftware/Hydra/lerna"
	"fmt"
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

	fmt.Println("in sessiondetails")	
	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.SessionAPI").GetString("uri")
	felicityusername := ConfObj.Sub("components.otrs.apis.SessionAPI.parameters").GetString("UserLogin")
	felicitypassword := ConfObj.Sub("components.otrs.apis.SessionAPI.parameters").GetString("Password")
        url := felicitybaseurl+felicityapiuri+"?UserLogin="+felicityusername+"&Password="+felicitypassword
       	fmt.Println(url)
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

