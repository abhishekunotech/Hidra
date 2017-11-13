package utils

import(
    "fmt"
	"github.com/Unotechsoftware/Hydra/lerna"
    "strings"
)

// Function that executes an action on the URL and URI passed to it
//
// Business Logic: Reads the action string, url and uri, builds a GET URL String
//
// Makes a call to writer of HTTP
func ExecuteCallGet(Component string, URI string, Action string) []uint8 {
        ConfObj := lerna.ReturnConfigObject()

        felicitybaseurl := ConfObj.Sub(Component).GetString("url")
        felicityapiuri := ConfObj.Sub(URI).GetString("uri")

        //READ THE URI Parameter ... e.g., components.otrs.apis.GetTicketBody
        // SEARCH THE 4th element in it... e.g., GetTicketBody
        // USE IT AS KEY in HMGET GetTicketBody ComponentURI
        // USE THE ComponentURI as felicityapiuri

        url := felicitybaseurl + felicityapiuri + Action
	return MakeHTTPGetCall(url)

}

func ExecuteCallGet1(Component string, URI string, Action string) []uint8{
    felicitybaseurl := "http://192.168.2.166/felicity/nph-genericinterface.pl/Webservice"
    felicityapiuri := strings.Split(URI,".")[3]
    fmt.Println(felicityapiuri+felicitybaseurl+Action)
    return ExecuteCallGet(Component,URI,Action)
}
