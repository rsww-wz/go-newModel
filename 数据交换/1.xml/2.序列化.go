package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

/*
func Marshal(v interface{}) ([]byte, error)
	参数：结构体/map(的地址)
	返回值：字节切片
	返回值是一个没有格式的xml文档，很难阅读
	所以有另外一个序列化函数，可以序列化成有格式的xml

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
	prefix   第二个参数为 xml 每行的前缀
	index    为标签嵌套时使用的空位符，这里选择用缩进值，也可以使用四个空格

序列化规则
	结构体的字段/map的key，作为xml的标签
	值作为xml的标签内的值

注意
	无论是 xml.Marshal 还是 xml.MarshalIndent 方法
	其传入的第一个参数值类型是 interface{} 类型，也就是用于进行序列化的对象
	传入的可以是对象，也可以是对象的地址
	但是建议传入地址
*/

//通过序列化我们可以将该对象变为一个 XML 字符串
func marshalStruct(v interface{}) string {
	// 没有格式
	//result, err := xml.Marshal(&info)

	// 格式化
	result, err := xml.MarshalIndent(&v, "", "\t")

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(result)
}

// xml 文件开头有一个头
// <?xml version="1.0"?>
// 这个头在序列化的结果中是没有的，所以得自己写入

func WriteXML1(address string, v interface{}) {
	data, err := xml.MarshalIndent(&v, "", "\t")

	if err != nil {
		fmt.Println("序列化出错")
	}

	data = append([]byte("<?xml version=\"1.0\"?>\n"), data...)

	// 写出文件
	err = ioutil.WriteFile(address, data, 0666)
	if err != nil {
		fmt.Println(err)
	}
}

func WriteXML2(address, data string) {
	head := "<?xml version=\"1.0\"?>\n"
	data = head + data

	file, err := os.OpenFile(address, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
