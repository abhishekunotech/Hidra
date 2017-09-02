package lerna

import(
	"fmt"
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


func GetJSONObjectType_Version(){


}


func GetJSONObjectType_Routes(){


}


func GetJSONObjectType_Components(){


}
