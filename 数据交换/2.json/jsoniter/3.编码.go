package jsonTest

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strings"
)

func Test() {
	jsonStr := `{"Name":"小刚","Gender":false,"Age":25,"Telephone":"13638562693","Book":{"Name":"水浒传","Author":"斯耐庵","Price":89.9}}`
	jsonByte := []byte(jsonStr)

	// 任何属性的类型都可以用ToString()
	// 无法直接获取嵌套内的字段属性，可以多次使用get函数
	// 当属性值是切片，数组时，可以在后面跟索引,索引从0开始
	//fmt.Println(jsoniter.Get(jsonByte, "Book",0).ToString())

	fmt.Println(jsoniter.Get(jsonByte, "Name").ToString())
	fmt.Println(jsoniter.Get(jsonByte, "Gender").ToString())
	fmt.Println(jsoniter.Get(jsonByte, "Age").ToString())
	fmt.Println(jsoniter.Get(jsonByte, "Book").ToString())

	fmt.Println(jsoniter.Get(jsonByte, "Book").Get("Name").ToString())
	fmt.Println(jsoniter.Get(jsonByte, "Book").Get("Author").ToString())
	fmt.Println(jsoniter.Get(jsonByte, "Book").Get("Price").ToString())
	fmt.Println()

	fmt.Println(jsoniter.Get(jsonByte).Keys())
	fmt.Println(jsoniter.Get(jsonByte, "Book").Keys())
	fmt.Println()

	// 属性
	fmt.Println("属性")
	fmt.Println(jsoniter.Get(jsonByte).Size())
	fmt.Println(jsoniter.Get(jsonByte).GetInterface())
	fmt.Println(jsoniter.Get(jsonByte).Get("Book").GetInterface())
	fmt.Println()

	// 反序列化
	reader := strings.NewReader(jsonStr)
	decoder := jsoniter.NewDecoder(reader)
	container := make(map[string]interface{})
	err := decoder.Decode(&container)
	if err !=nil {
		fmt.Println(err)
		return
	}

	fmt.Println(container)

	// 序列化
	var buf bytes.Buffer
	// 把内容写入buf
	encode := jsoniter.NewEncoder(&buf)

	// 设置格式，不支持转义字符
	encode.SetIndent("","    ")

	// 进行json编码
	if err := encode.Encode(container);err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf.Bytes()))
}
