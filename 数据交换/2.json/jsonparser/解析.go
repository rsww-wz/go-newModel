package jsonparser

import (
	"fmt"
	"github.com/buger/jsonparser"
)

var data = []byte(`{
    "Name": "小刚",
    "Gender": false,
    "Age": 25,
    "Telephone": "13638562693",
    "Book": {
        "Name": "水浒传",
        "Author": "斯耐庵",
        "Price": 89.9
    }
}`)

// 获取类型
func GetType() {
	r1, t1, size1, err := jsonparser.Get(data)
	r2, t2, size2, err := jsonparser.Get(data, "Gender")
	r3, t3, size3, err := jsonparser.Get(data, "Age")
	r4, t4, size4, err := jsonparser.Get(data, "Name")
	r5, t5, size5, err := jsonparser.Get(data, "Book")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("value = %s, type = %v, size = %d\n", r1, t1, size1)
	fmt.Printf("value = %s, type = %v, size = %d\n", r2, t2, size2)
	fmt.Printf("value = %s, type = %v, size = %d\n", r3, t3, size3)
	fmt.Printf("value = %s, type = %v, size = %d\n", r4, t4, size4)
	fmt.Printf("value = %s, type = %v, size = %d\n\n\n", r5, t5, size5)
}

// 获取值
func GetValue() {
	r1, err := jsonparser.GetString(data, "Name")
	r2, err := jsonparser.GetInt(data, "Age")
	r3, err := jsonparser.GetBoolean(data, "Gender")
	r4, err := jsonparser.GetFloat(data, "Book", "Price")

	// 字符串类型的数字
	r5, err := jsonparser.GetInt(data, "Telephone")    // 0
	r6, err := jsonparser.GetString(data, "Telephone") // 13638562693

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println(r5)
	fmt.Println(r6)
}

// 解析成原来的值，整型解析成int64，浮点型解析成float64
// 直接输入要解析的字节切片
func ParseVal() {

}

// 修改值
func Modify() {
	v1, err := jsonparser.Set(data, []byte("22"), "Age")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(v1))

	// 修改完之后，再看看类型有没有被改变
	r, _ := jsonparser.GetInt(v1, "Age")
	fmt.Println(r)

	// 修改成与原value不符的类型,再获取它的类型
	v2, _ := jsonparser.Set(data, []byte("男"), "Gender")
	fmt.Println(string(v2))                      // 修改成功
	r1, _ := jsonparser.GetBoolean(v2, "Gender") // false
	r2, _ := jsonparser.GetString(v2, "Gender")  // 空
	fmt.Println(r1)
	fmt.Println(r2)
}

// 删除key，value
func Remove() {
	// 如果key不存在，则不删除
	v1 := jsonparser.Delete(data, "Gender")
	fmt.Println(string(v1))
}
