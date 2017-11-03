package utils

import(
	"github.com/Unotechsoftware/Hydrav2/lerna"
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
        url := felicitybaseurl + felicityapiuri + Action
	return MakeHTTPGetCall(url)

}


