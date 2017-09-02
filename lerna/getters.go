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

func GetComponentType_Version(){

}

func GetComponentType_Url(){

}

func GetComponentType_Apis(){

}

func GetAPIType_Version(){


}


func GetAPIType_Uri(){

}

func GetAPIType_Parameter(){


}


func GetRouteType_Version(){

}


func GetRouteType_Method(){

}

func GetRouteType_Uri(){


}


func GetRouteType_Component(){



}


func GetRouteType_Handler(){


}


func GetJSONObjectType_Version(abc *viper.Viper) string{
	return	abc.GetString("version")
}


func GetJSONObjectType_Routes(abc *viper.Viper) []RouteType{
	returnVal := abc.GetStringMap("routes")
	keys_of_returnval := GetKeyArray(returnVal)
	RouteInside := abc.Sub("routes")
	//var returnStrg string
	var returnValue []RouteType

	// Need to define the size of the array. 
	// Or append element to slice
	for _, element := range keys_of_returnval{
		var Meow RouteType
		_ = RouteInside.UnmarshalKey(element, &Meow)	
		returnValue = append(returnValue,Meow)
	}

	return returnValue
}


func GetJSONObjectType_Components(){


}
