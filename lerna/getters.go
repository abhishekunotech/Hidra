package lerna

import(
	"fmt"
	"github.com/spf13/viper"
	//"reflect"
)

func GetKeyArray(Abc map[string]interface{}) []string{
        keys := make([]string, len(Abc))

        i := 0
        for k := range Abc {
                keys[i] = k
                i++
        }

        fmt.Println("\n\n\n\n KEYS ARRAY \n\n\n\n")

        fmt.Println(keys)

        fmt.Println("\n\n\n End of Keys Array \n\n\n\n")

        return keys

}

func GetComponentType_Version(abc ComponentType) string{
	return abc.Version
}

func GetComponentType_Url(abc ComponentType) string{
	return abc.Url
}

func GetComponentType_Apis(abc ComponentType) map[string]APIType{
	return abc.Apis
}

func GetAPIType_Version(abc APIType) string{
	return abc.Version
}


func GetAPIType_Uri(abc APIType) string{
	return abc.Uri
}

func GetAPIType_ParameterObj(abc APIType) map[string]ParameterVal{
	return abc.Parameter
}

func GetAPIType_Parameter(abc ParameterVal) string{
	return abc.Value
}

func GetRouteType_Version(abc RouteType) string{
	return abc.Version
}


func GetRouteType_Method(abc RouteType) string{
	return abc.Method
}

func GetRouteType_Uri(abc RouteType) string{
	return abc.Uri
}


func GetRouteType_Component(abc RouteType) string{
	return abc.Component
}


func GetRouteType_Handler(abc RouteType) string{
	return abc.Handler
}


func GetJSONObjectType_Version(abc *viper.Viper) string{
	return	abc.GetString("version")
}


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
