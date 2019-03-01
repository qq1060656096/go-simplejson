package simplejson

import (
"encoding/json"
"errors"
"reflect"
"strconv"
)

// Map type 断言json对象
func (j *Json) Object() (map[string]interface{}, error) {
	if m, ok := (j.data).(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("type assertion to map[string]interface{} failed")
}

// Array type 断言json数组
func (j *Json) Array() ([]interface{}, error) {
	if arr, ok := (j.data).([]interface{}); ok {
		return arr, nil
	}
	return nil, errors.New("type assertion to []interface{} failed")
}

// String type 断言json字符串值
func (j *Json) String() (string, error) {
	if s, ok := (j.data).(string); ok {
		return s, nil
	}
	t := reflect.TypeOf(j.data)
	t2:= reflect.TypeOf(json.Number("1"))
	if t == t2 {
		return string(j.data.(json.Number)), nil
	}
	return "", errors.New("type assertion to string failed")
}

// Bool type 断言json布尔值
func (j *Json) Bool() (bool, error) {
	if b, ok := (j.data).(bool); ok {
		return b, nil
	}
	return false, errors.New("type assertion to string failed")
}
// Int type 断言json int值
func (j *Json) Int() (int, error) {
	switch j.data.(type) {
	case json.Number:
		i, err := j.data.(json.Number).Int64()
		return int(i), err
	case float32, float64:
		return int(reflect.ValueOf(j.data).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(j.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(j.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

// Int64 type 断言json int64值
func (j *Json) Int64() (int64, error) {
	switch j.data.(type) {
	case json.Number:
		return j.data.(json.Number).Int64()
	case float32, float64:
		return int64(reflect.ValueOf(j.data).Float()), nil
	case int, int8, int16, int32, int64:
		return int64(reflect.ValueOf(j.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(j.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

// Uint64 type 断言json uint64 值
func (j *Json) Uint64() (uint64, error) {
	switch j.data.(type) {
	case json.Number:
		return strconv.ParseUint(j.data.(json.Number).String(), 10, 64)
	case float32, float64:
		return uint64(reflect.ValueOf(j.data).Float()), nil
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(j.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(j.data).Uint(), nil
	}
	return 0, errors.New("invalid value type")
}

// Float32 type 断言json float32 值
func (j *Json) Float32() (float32, error) {
	v, err := j.Float64()
	if err != nil {
		return 0, errors.New("invalid value type")
	}
	return float32(v), nil
}

// Float64 type 断言json float64 值
func (j *Json) Float64() (float64, error) {
	switch j.data.(type) {
	case json.Number:
		v, err := j.data.(json.Number).Float64()
		return v, err
	case float32, float64:
		return float64(reflect.ValueOf(j.data).Float()), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(j.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(j.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}