package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

/*
func Unmarshal(data []byte, v interface{}) error
	反序列化的函数 xml.Unmarshal 和序列化的函数不同
	其传入的第二个参数必须是用于接受结果的结构体对象的地址

反序列化过程中不需要专门处理 XML 文件头
Go 反序列化过程会自动忽略 XML 文件头
*/

func unMarshalStruct(address string, v interface{}){
	data,err := ioutil.ReadFile(address)
	if err != nil {
		fmt.Println("read file fail...")
		return
	}

	// 反序列化
	err = xml.Unmarshal(data,v)
	if err != nil {
		fmt.Println("反序列化失败")
	}

}