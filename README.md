### go-simplejson

go语言操作json包

### 文档

### 1. 字符串解析成json, 并获取值
```go
j, err := simplejson.NewJson([]byte(jsonStr))
j.Get("name").String()
j.Get("profile").Get("age").Int()
j.Get("profile").Get("mobile").GetArrayIndex(1).Int()
```
####  示例
```go
package main

import (
	"fmt"
	"github.com/qq1060656096/go-simplejson"
)

func main() {
	jsonStr := `
{
	"uid": 1,
	"name": "tester1",
	"pass": "123456",
	"profile": {
		"age": 18,
		"weight": "75kg",
		"height": "1.71m",
		"mobile": [
			15400000001,
			15400000002
		]
	}
}
`
	// 字符串解析成json对象
	j, err := simplejson.NewJson([]byte(jsonStr))

	// 简单获取值, 并转换成string类型
	nameValue, err := j.Get("name").String()
	fmt.Println(err)
	fmt.Println(nameValue) // 输出: tester1

	// 连贯操作获取子级键的值, 并转换成int类型
	ageValue, err := j.Get("profile").Get("age").Int()
	fmt.Println(ageValue) // 输出: 18

	// 连贯操作获取子级数组索引值, 并转换成int类型
	mobileIndex2Value, err := j.Get("profile").Get("mobile").GetArrayIndex(1).Int()
	fmt.Println(mobileIndex2Value) // 输出: 15400000002
}
```
### json类型对应golang类型
```go
boolean >> bool
number  >> float64
string  >> string
null    >> nil
array   >> []interface{}
object  >> map[string]interface{}
```