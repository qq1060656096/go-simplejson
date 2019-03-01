### go-simplejson

go语言操作json包

[![Build Status](https://secure.travis-ci.org/bitly/go-simplejson.png)](http://travis-ci.org/bitly/go-simplejson)

### 导入包

    import github.com/qq1060656096/go-simplejson

### 文档

### 解析json字符串
```go
import (
	"github.com/qq1060656096/go-simplejson"
	"fmt"
)

body := `
{
	"id": 10,
	"name": "test4",
	"age": 18,
	"address": {
		"country": "中国",
		"city": "成都"
	},
	"mobile": [
		"15400012301",
		"15400012302",
		15400012303
	]
}
`
// json字符解析
j, err := NewJson([]byte(body))
// 获取json对象值
v, err := j.Get("name").String()

// 获取json对象值,连贯操作
v, err := j.Get("address").Get("city").String()
fmt.printf("%s", v)// 成都

// 获取json对象值
v, err := j.Get("name").String()
fmt.printf("%s", v)// test4

// 获取json数组索引值
v, err := j.Get("mobile").GetArrayIndex(2).Int()
fmt.printf("%s", v)// 15400012303
```