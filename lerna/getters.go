package lerna

import(
	"github.com/spf13/viper"
)

// Function that returns the list of all the keys defined in the configuration
//
// Business Logic: Provides a List of Keys to make a call to Get the value of that key
//
// Returns Key Array as a string slice
func GetKeyArray(Abc map[string]interface{}) []string{
        keys := make([]string, len(Abc))

        i := 0
        for k := range Abc {
                keys[i] = k
                i++
        }


        return keys

}


// Function that returns the Version of a Component
// 
// Business Logic: Gets the Value corresponding to the key "Version" in Component Object
//
// Returns a Version String
func GetComponentType_Version(abc ComponentType) string{
	return abc.Version
}

// Function that returns the URL of a Component
//
// Business Logic: Gets the Value corresponding to the key "URL" in Component Object
//
// Returns a URL String
func GetComponentType_Url(abc ComponentType) string{
	return abc.Url
}

// Function that returns the APIs of a Component
//
// Business Logic: Gets the Value corresponding to the key "APIs" in Component Object
//
// Returns a map[string] of type API
func GetComponentType_Apis(abc ComponentType) map[string]APIType{
	return abc.Apis
}

// Function that returns the Version of an API
//
// Business Logic: Gets the Value corresponding to the key "Version" in API Object
//
// Returns a Version String
func GetAPIType_Version(abc APIType) string{
	return abc.Version
}

// Function that returns the URI corresponding to API Type
//
// Business Logic: Gets the Value corresponding to the key "URI" in API Object
//
// Returns a URI String
func GetAPIType_Uri(abc APIType) string{
	return abc.Uri
}

// Function that returns the Parameter of APIType
//
// Business Logic: Gets the Value corresponding to the key "ParameterObj" in API Object
//
// Returns a map[string] of type of ParameterVal Type
func GetAPIType_ParameterObj(abc APIType) map[string]ParameterVal{
	return abc.Parameter
}

// Function that returns the Parameter of a API
//
// Business Logic: Gets the Value corresponding to the key "Parameter" in API Object
//
// Returns a Parameter String
func GetAPIType_Parameter(abc ParameterVal) string{
	return abc.Value
}


// Function that returns the Version of a Route
//
// Business Logic: Gets the Value corresponding to the key "Version" in Route Object
//
// Returns a Version String
func GetRouteType_Version(abc RouteType) string{
	return abc.Version
}

// Function that returns the Method of a Route
//
// Business Logic: Gets the Value corresponding to the key "Method" in Route Object
//
// Returns a Method String
func GetRouteType_Method(abc RouteType) string{
	return abc.Method
}

// Function that returns the URI of a Route
//
// Business Logic: Gets the Value corresponding to the key "URI" in Route Object
//
// Returns a URI String
func GetRouteType_Uri(abc RouteType) string{
	return abc.Uri
}

// Function that returns the Component of Route
//
// Business Logic: Gets the Value corresponding to the key "Component" in Route Object
//
// Returns a Component String
func GetRouteType_Component(abc RouteType) string{
	return abc.Component
}

// Function that returns the Version of a Handler in Routes
//
// Business Logic: Gets the Value corresponding to the key "Version" in Route Object's Handler
//
// Returns a Version String
func GetRouteType_Handler(abc RouteType) string{
	return abc.Handler
}

// Function that returns the Version of the Configuration
//
// Business Logic: Gets the Value corresponding to the key "Version" of the Configuration
//
// Returns a Version String
func GetJSONObjectType_Version(abc *viper.Viper) string{
	return	abc.GetString("version")
}

// Function that returns a HashMap with keys and values of Route Object
//
// Business Logic: Gets the List of Key-Value Pairs corresponding to the Object "Route" in Hydra Configuration
//
// Returns a map[string] of Type RouteType
func GetJSONObjectType_Routes(abc *viper.Viper) map[string]RouteType{
	returnVal := abc.GetStringMap("routes")
	keys_of_returnval := GetKeyArray(returnVal)
	RouteInside := abc.Sub("routes")
	returnValue := make(map[string]RouteType)

	for _, element := range keys_of_returnval{
		var Element RouteType
		_ = RouteInside.UnmarshalKey(element, &Element)	
		returnValue[element] = Element
	}

	return returnValue
}

// Function that returns a HashMap with keys and values of Component Object
//
// Business Logic: Gets the List of Key-Value Pairs corresponding to the Object "Component" in Hydra Configuration
//
// Returns a map[string] of Type ComponentType
func GetJSONObjectType_Components(abc *viper.Viper) map[string]ComponentType{
	returnVal := abc.GetStringMap("components")
	

	keys_of_returnval := GetKeyArray(returnVal)
	ComponentInside := abc.Sub("components")

	returnValue := make(map[string]ComponentType)

	for _,element := range keys_of_returnval{
		var Element ComponentType
		fasfa := ComponentInside.Sub(element)
		_ = fasfa.Unmarshal(&Element)
		returnValue[element] = Element
		
	}
	return returnValue
}
