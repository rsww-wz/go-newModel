package jsonTest

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

// 也是可以忽略json中的格式化
func JsoniterUnMarsh(v interface{}, data []byte) {
	err := jsoniter.Unmarshal(data, v)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UnMarshJsonStr(v interface{}, str string) {
	if err := jsoniter.UnmarshalFromString(str, v); err != nil {
		fmt.Println(err)
	}
}
