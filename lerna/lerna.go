package lerna

import (
	"github.com/antigloss/go/logger"
	"github.com/spf13/viper"
)

// Json Object Type Returns Level - 1 structure of Hydra Configuration
type JSONObjectType struct {
	Version    string `json:"version"`
	Routes     map[string] RouteType
	Components map[string] ComponentType
}

// Route Object returns array elements from Route Object in JSON - Hydra Configuration
type RouteType struct {
	Version   string `json:"version"`
	Method    string `json:"method"`
	Uri       string `json:"URI"`
	Component string `json:"component"`
	Handler   string `json:"handler"`
}

// Component Object returns array elements from Component Object in JSON - Hydra Configuration
type ComponentType struct {
	Version string `json:"version"`
	Url     string `json:"URL"`
	Apis    map[string] APIType
}

// API Object returns array elements from API Type Object in JSON - Hydra Configuration
type APIType struct {
	Version   string `json:"version"`
	Uri       string `json:"URI"`
	Parameter map[string] ParameterVal
}

// Returns Values of Parameters
type ParameterVal struct {
	Value string 
}



// Returns a Viper Object to access the Configuration
func ReturnConfigObject() *viper.Viper{

	ViConfig := viper.New()
	ViConfig.SetConfigName("Hydra")
	ViConfig.AddConfigPath("/etc/Hydra/conf.d/")
	ViConfig.SetConfigType("json")

	err := ViConfig.ReadInConfig() // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		logger.Error(err.Error())
	}

	
	
	return ViConfig

}
