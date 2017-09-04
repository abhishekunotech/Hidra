package main

import (
        "fmt"
	"io/ioutil"
	"encoding/json"
        //"github.com/spf13/viper"
)

type JSONObjectType struct {

	Version    string `json:"version"`

	Routes     []RouteType

	Components []ComponentType

}



type RouteType struct {

	Version   string `json:"version"`

	Method    string `json:"method"`

	Uri       string `json:"URI"`

	Component string `json:"component"`

	Handler   string `json:"handler"`

}



type ComponentType struct {

	Version string `json:"version"`

	Url     string `json:"URL"`

	Apis    []APIType

}



type APIType struct {

	Version   string `json:"version"`

	Uri       string `json:"URI"`

	Parameter map[string]string

}

func main() {

	/*viper.SetConfigType("json")
        viper.SetConfigFile("/etc/Hydra/conf.d/Hydra.js")
	//vi := New()*Viper
	//var vi *Viper
	err := viper.ReadInConfig()
 
        if err != nil {
                fmt.Println("No configuration file")
        } 

	viper.WatchConfig()

	fmt.Println("Hello")

        var NewVar2 interface{}

        _ = viper.Unmarshal(&NewVar2)

        fmt.Println(NewVar2)
	fmt.Println("Bye")

	vc := viper.Get("routes")

	fmt.Println(vc)
     
	co := viper.GetString("index.method")
	fmt.Println(co)

	components := viper.Get("components")

	fmt.Println(components)
	*/

	file, e := ioutil.ReadFile("/etc/Hydra/conf.d/Hydra.js")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
    }


var p JSONObjectType
if err := json.Unmarshal([]byte(JSONObjectType), &p); err != nil {
    fmt.Println(err)
}
for k, node := range p.Routes {
    fmt.Printf("%s: %s\n", k, node.Version )
}



}












