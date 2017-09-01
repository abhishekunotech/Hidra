package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"path/filepath"
	//"encoding/json"
	//"bytes"
)

type config struct {
	configPaths    []string
	configName     string
	configFile     string
	configType     string
	onConfigChange func(fsnotify.Event)
}

type jsonobject struct {
	Version    string `json:"version"`
	Routes     []RoutesType
	Components []ComponentsType
}

type RoutesType struct {
	Name    string `json:"name"`
	Method  string `json:"method"`
	URI     string `json:"URI"`
	Handler string `json:"handler"`
}

type ComponentsType struct {
	ComponentName string `json:"componentName"`
	URL           string `json:"url"`
	API           []APIs
}

type APIs struct {
	Name       string `json:"name"`
	URI        string `json:"URI"`
	Parameters []Params
}

type Params struct {
	TicketId  string `json:"TicketId"`
	UserLogin string `json:"UserLogin"`
	Password  string `json:"Password"`
}

func main() {

	file, e := ioutil.ReadFile("/etc/Hydra/conf.d/Hydra.js")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
	}
}
