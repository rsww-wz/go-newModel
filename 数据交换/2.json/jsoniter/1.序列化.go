package jsonTest

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)


func JsoniterMarsh(data interface{}) {

	b, err := jsoniter.Marshal(data)
	bb, err :=  jsoniter.MarshalIndent(data, "", "    ")
	str,err := jsoniter.MarshalToString(data)
	if err != nil{
		fmt.Println("error: ", err)
	}

	fmt.Println(string(b))
	fmt.Println()
	fmt.Println(str)
	fmt.Println()
	fmt.Println(string(bb))
}