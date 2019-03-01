package simplejson

import (
	"bytes"
	"encoding/json"
)


// DecodeJSON 字符串解析json
func (j *Json) DecodeJSON(b []byte) error {
	r := bytes.NewBuffer(b)
	d := json.NewDecoder(r)
	d.UseNumber()
	return d.Decode(&j.data)
}

// EncodeJSON Json对象转换成字节
func (j *Json) EncodeJSON() ([]byte, error) {
	b, err := json.Marshal(&j.data)
	return b, err
}

// EncodeJsonPretty Json对象转换成字节
func (j *Json) EncodeJsonPretty() ([]byte, error) {
	return json.MarshalIndent(&j.data, "", "  ")
}