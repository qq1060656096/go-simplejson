package simplejson

import "log"


// Json 操作json对象
type Json struct {
	data interface{}
}

// NewJson returns 返回一个指针指向新的Json对象
func NewJson(data []byte) (*Json, error) {
	j := new(Json)
	err := j.DecodeJSON(data)
	return j, err
}

// New returns 一个指针指向新的JsonK空对象
func New() *Json {
	j := new(Json)
	j.data = make(map[string]interface{})
	return j
}

// Get 获取json 对象字段值
func (j *Json) Get(key string) *Json {
	m, err := j.Object()
	if err == nil {
		if v, ok := m[key]; ok {
			return &Json{v}
		}
	}
	return &Json{nil}
}

// GetArrayIndex type 索引值
func (j *Json) GetArrayIndex(index int) *Json {
	v, err := j.Array()
	switch {
	case err == nil && len(v) > 0:
		return &Json{v[index]}
	}
	return &Json{nil}
}

// GetInterface 直接获取当前值
func (j *Json) GetInterface() interface{} {
	return j.data
}

// MustSet 设置数据
// keys type 必须是string或者int 否则报错
func (j *Json) MustSet(value interface{}, keys ...interface{}) {
	j.data = mustSetData(j.data, value, keys...)
}

// mustSetData 设置数据
func mustSetData(data interface{}, value interface{}, keys ...interface{}) interface{} {
// 定义函数: mustSetData (data interface{}, value interface{}, keys ...interface{}) interface{}
// 定义变量: data 要设置的数据, value设置的值, keys设置的一组键
// 1. 获取当前keys的长度, keysLen = 当前keys的长度
// 2. 获取当前keys的第0个索引值, currkey = keys[0]
// 3. 断言currKey的类型，如果不是int, string就跑出致命错误
// 4. 如果 currKey 是字符串类型
//		4.1 如果data == nil 就初始化data = make(map[string]interface{}), 防止nil转换map[string]interface{}错误
//		4.2 把currKey转换成string类型
// 		4.3 把data转换成map[string]interface{}类型
//		4.4 获取当前key的值(用于多层级key递归传到下一次处理), currValue, ok = data[currKey]
// 		4.5 如果 keysLen > 1 代表需要递归处理
//			4.5.1 currKey 不存在, 就把 currValue = nil
// 			4.5.2 调用 mustSetData递归处理, 传递当前键的currValue值，要设置的value值和除0外的一组键keys[1:],
// 					然后把返回值设置为当前键的currValue值, currValue = mustSetData(currValue = 当前键的值, value 要设置的值, keys[1:]... = 除第0个索引外的一组键)
//			4.5.3 设置数据键的值, 并返回data, data[currKey] = currValue, return data
// 		4.6 如果是最后一个键，就直接设置data值, 并返回data数据, data[currKey] = value, return data
// 5. 如果 currKey 是整型
// 	 	5.1 如果data == nil, 就初始化 data = make([]interface{}, 0)
// 		5.2 把currKey转换成 int类型
// 		5.3 把data转换成 []interface{}类型
// 		5.4 如果 currIndex在 data索引范围内就把currValue设置为 currValue = data[currIndex]
// 		5.5 如果 keysLen > 1 代表需要递归处理
//				5.5.1 调用 mustSetData递归处理, 传递当前键的currValue值，要设置的value值和除0外的一组键keys[1:],
//	 					然后把返回值设置为当前键的currValue值, currValue = mustSetData(currValue = 当前键的值, value 要设置的值, keys[1:]... = 除第0个索引外的一组键)
// 		5.6 如果是最后一个键，就直接设置data值, 并返回data数据, data[currIndex] = value, return data
	keysLen := len(keys)
	currKey := keys[0]
	switch currKey.(type) {
	case string:
		if (data == nil) {
			data = make(map[string]interface{})
		}
		currKey, _ := currKey.(string)
		//转换data类型
		data := data.(map[string]interface{})
		// 获取键的值
		currValue, ok := data[currKey]
		// 如果键keysLen长度 > 1就递归处理
		if keysLen > 1 {
			// 键不存在
			if !ok {
				currValue = nil
			}
			currValue = mustSetData(currValue, value, keys[1:]...)
			return setObject(data, currKey, currValue)
		}
		return setObject(data, currKey, value)
	case int:
		if (data == nil) {
			data = make([]interface{}, 0)
		}
		currIndex, _ := currKey.(int)
		// 转换data类型
		data := data.([]interface{})
		arrLen := len(data)
		// 获取键的值
		var currValue interface {}
		// 索引存在
		if currIndex >= 0 && currIndex < arrLen {
			currValue = data[currIndex]
		}
		// 如果键keysLen长度 > 1就递归处理
		if keysLen > 1 {
			currValue = mustSetData(currValue, value, keys[1:]...)
			return setArray(data, currIndex, currValue)
		}
		return setArray(data, currIndex, value)
	}
	log.Panicf("simplejson.SetData keys type must string or int, key type is %T", currKey)
	return nil
}

// setObject 设置对象值
func setObject(obj map[string]interface{}, key string, value interface{}) interface{} {
	if obj == nil {
		obj = make(map[string]interface{})
	}
	obj[key] = value
	return obj
}

// setArray 设置数组值
func setArray(arr []interface{}, index int, value interface{}) (interface{}) {
	if arr == nil {
		arr = make([]interface{}, 0)
	}
	arrLen := len(arr)
	if index < 0 || index >= arrLen {
		arr = append(arr, value)
		return arr
	}
	arr[index] = value
	return arr
}