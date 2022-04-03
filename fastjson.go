package fastjson

import (
	"github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// JSONArray 数组类型的json
type JSONArray interface {
	// Size 获取数量
	Size() int
	// IsZero 判断零值
	IsZero() bool
	// ContainsKey 是否存在索引index
	ContainsIndex(index int) bool
	// GetInt 获取int类型的值
	GetInt(index int) (int, error)
	// GetInt32 获取int32类型的值
	GetInt32(index int) (int32, error)
	// GetInt64 获取int64类型的值
	GetInt64(index int) (int64, error)
	// GetFloat32 获取float32类型的值
	GetFloat32(index int) (float32, error)
	// GetFloat64 获取float64类型的值
	GetFloat64(index int) (float64, error)
	// GetString 获取string类型的值
	GetString(index int) (string, error)
	// GetBool 获取bool类型的值
	GetBool(index int) (bool, error)
	// GetJSONObject 获取JSONObject类型的值
	GetJSONObject(index int) (JSONObject, error)
	// GetJSONArray 获取JSONArray类型的值
	GetJSONArray(index int) (JSONArray, error)
	// GetValue 获取原始值
	GetValue() ([]interface{}, error)
	// Put 将一个值加入JSONArray
	Put(value interface{}) error
	// MarshalJSON 实现标准库json接口
	MarshalJSON() ([]byte, error)
	// Scan 转为其他类型的值
	Scan(dest []interface{}) error
}

// JSONOject 对象类型的json
type JSONObject interface {
	ToJSONString() string
	// IsZero 获取int类型值
	IsZero() bool
	// ContainsKey 是否存在键key
	ContainsKey(key string) bool
	// GetInt 获取int类型的值
	GetInt(key string) (int, error)
	// GetInt32 获取int32类型的值
	GetInt32(key string) (int32, error)
	// GetInt64 获取int64类型的值
	GetInt64(key string) (int64, error)
	// GetFloat32 获取float32类型的值
	GetFloat32(key string) (float32, error)
	// GetFloat64 获取float64类型的值
	GetFloat64(key string) (float64, error)
	// GetString 获取string类型的值
	GetString(key string) (string, error)
	// GetBool 获取bool类型的值
	GetBool(key string) (bool, error)
	// GetJSONObject 获取JSONObject类型的值
	GetJSONObject(key string) (JSONObject, error)
	// GetJSONArray 获取JSONArray类型的值
	GetJSONArray(key string) (JSONArray, error)
	// GetValue 获取原始值
	GetValue() (map[string]interface{}, error)
	// Put 将一个值加入JSONArray
	Put(key string, value interface{}) error
	// MarshalJSON 实现标准库json接口
	MarshalJSON() ([]byte, error)
	// Scan 转为其他类型的值
	Scan(dest interface{}) error
}

type jsonObject struct {
	value map[string]interface{}
}

type jsonArray struct {
	value []interface{}
}

// NewJSONObjectFrom 构造JSONObject
func NewJSONObjectFrom(jsonByte []byte) (JSONObject, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(jsonByte, &data); err != nil {
		return nil, err
	}
	return &jsonObject{
		value: data,
	}, nil
}

func NewJSONObject() JSONObject {
	var data = make(map[string]interface{})
	return &jsonObject{
		value: data,
	}
}

// NewJSONArrayFrom 构造JSONArray
func NewJSONArrayFrom(jsonByte []byte) (JSONArray, error) {
	var data []interface{}
	if err := json.Unmarshal(jsonByte, &data); err != nil {
		return nil, err
	}
	return &jsonArray{
		value: data,
	}, nil
}

func NewJSONArray() JSONArray {
	var data = make([]interface{}, 0)
	return &jsonArray{
		value: data,
	}
}

func ToJSONString(v interface{}) string {
	b, _ := json.MarshalToString(v)
	return b
}

func ToJSONBytes(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func ToJSON(v interface{}) (obj JSONObject) {
	if b, err := json.Marshal(v); err != nil {
		return nil
	} else {
		obj = new(jsonObject)
		err = json.Unmarshal(b, obj)
		if err != nil {
			return nil
		}
	}
	return obj
}
func ParseObject(b []byte) (JSONObject, error) {
	var value map[string]interface{}
	err := json.Unmarshal(b, &value)
	if err != nil {
		return nil, err
	}
	return &jsonObject{value: value}, nil
}

func ParseObjectV2[T any](b []byte, v T) (*T, error) {
	err := json.Unmarshal(b, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
