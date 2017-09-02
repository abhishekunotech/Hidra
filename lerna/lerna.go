package lerna

import (
	//"fmt"
	"github.com/antigloss/go/logger"
	"github.com/spf13/viper"
)


type JSONObjectType struct {
	Version    string `json:"version"`
	Routes     map[string] RouteType
	Components map[string] ComponentType
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
	Apis    map[string] APIType
}

type APIType struct {
	Version   string `json:"version"`
	Uri       string `json:"URI"`
	Parameter map[string] ParameterVal
}

type ParameterVal struct {
	Value string 
}


//Write Initialize Redis in Utils
//Import and Call that MOFO here
func ReturnConfigObject() *viper.Viper{

	ViConfig := viper.New()
	ViConfig.SetConfigName("Hydra_non")
	ViConfig.AddConfigPath("/etc/Hydra/conf.d/")
	ViConfig.SetConfigType("json")

	err := ViConfig.ReadInConfig() // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		logger.Error(err.Error())
	}

	
	
	return ViConfig

}
