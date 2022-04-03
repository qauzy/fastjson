package fastjson

import (
	"fmt"
	"strconv"
)

func (obj *jsonObject) IsZero() bool {
	return obj.value == nil
}

func (obj *jsonObject) ContainsKey(key string) bool {
	_, ok := obj.value[key]
	return ok
}

func (obj *jsonObject) ToJSONString() string {
	if obj.value == nil {
		return ""
	}
	return ToJSONString(obj.value)
}

func (obj *jsonObject) GetInt(key string) (int, error) {
	if !obj.ContainsKey(key) {
		return 0, fmt.Errorf("key %s not exsits", key)
	}
	v := obj.value[key]
	switch t := v.(type) {
	case string:
		return strconv.Atoi(t)
	case []byte:
		return strconv.Atoi(string(t))
	case int:
		return t, nil
	case int32:
		return int(t), nil
	case int64:
		return int(t), nil
	case float32:
		return int(t), nil
	case float64:
		return int(t), nil
	default:
		return 0, fmt.Errorf("data is %T, not int", v)
	}
}

func (obj *jsonObject) GetInt32(key string) (int32, error) {
	if !obj.ContainsKey(key) {
		return 0, fmt.Errorf("key %s not exsits", key)
	}
	v := obj.value[key]
	switch t := v.(type) {
	case string:
		i, err := strconv.ParseInt(t, 10, 32)
		if err != nil {
			return 0, err
		}
		return int32(i), nil
	case []byte:
		i, err := strconv.ParseInt(string(t), 10, 32)
		if err != nil {
			return 0, err
		}
		return int32(i), nil
	case int:
		return int32(t), nil
	case int32:
		return t, nil
	case int64:
		return int32(t), nil
	case float32:
		return int32(t), nil
	case float64:
		return int32(t), nil
	default:
		return 0, fmt.Errorf("data is %T, not int32", v)
	}
}

func (obj *jsonObject) GetInt64(key string) (int64, error) {
	if !obj.ContainsKey(key) {
		return 0, fmt.Errorf("key %s not exsits", key)
	}
	v := obj.value[key]
	switch t := v.(type) {
	case string:
		return strconv.ParseInt(t, 10, 64)
	case []byte:
		return strconv.ParseInt(string(t), 10, 64)
	case int:
		return int64(t), nil
	case int32:
		return int64(t), nil
	case int64:
		return t, nil
	case float32:
		return int64(t), nil
	case float64:
		return int64(t), nil
	default:
		return 0, fmt.Errorf("data is %T, not int64", v)
	}
}

func (obj *jsonObject) GetFloat32(key string) (float32, error) {
	if !obj.ContainsKey(key) {
		return 0, fmt.Errorf("key %s not exsits", key)
	}
	v := obj.value[key]
	switch t := v.(type) {
	case string:
		f, err := strconv.ParseFloat(t, 32)
		if err != nil {
			return 0, err
		}
		return float32(f), nil
	case []byte:
		f, err := strconv.ParseFloat(string(t), 32)
		if err != nil {
			return 0, err
		}
		return float32(f), nil
	case int:
		return float32(t), nil
	case int32:
		return float32(t), nil
	case int64:
		return float32(t), nil
	case float32:
		return t, nil
	case float64:
		return float32(t), nil
	default:
		return 0, fmt.Errorf("data is %T, not float32", v)
	}
}

func (obj *jsonObject) GetFloat64(key string) (float64, error) {
	if !obj.ContainsKey(key) {
		return 0, fmt.Errorf("key %s not exsits", key)
	}
	v := obj.value[key]
	switch t := v.(type) {
	case string:
		return strconv.ParseFloat(t, 64)
	case []byte:
		return strconv.ParseFloat(string(t), 64)
	case int:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return t, nil
	default:
		return 0, fmt.Errorf("data is %T, not float64", v)
	}
}

func (obj *jsonObject) GetString(key string) (string, error) {
	if !obj.ContainsKey(key) {
		return "", fmt.Errorf("key %s not exsits", key)
	}
	v := obj.value[key]
	switch t := v.(type) {
	case string:
		return t, nil
	case []byte:
		return string(t), nil
	case int, int32, int64:
		return fmt.Sprint(t), nil
	case float32, float64:
		return fmt.Sprint(t), nil
	default:
		return "", fmt.Errorf("data is %T, not string", v)
	}
}

func (obj *jsonObject) GetBool(key string) (bool, error) {
	if !obj.ContainsKey(key) {
		return false, fmt.Errorf("key %s not exsits", key)
	}
	v := obj.value[key]
	t, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf("data is %T, not bool", v)
	}
	return t, nil
}

func (obj *jsonObject) GetJSONObject(key string) (JSONObject, error) {
	if !obj.ContainsKey(key) {
		return nil, fmt.Errorf("key %s not exsits", key)
	}
	v := obj.value[key]
	switch t := v.(type) {
	case string:
		return NewJSONObjectFrom([]byte(t))
	case []byte:
		return NewJSONObjectFrom(t)
	case JSONObject:
		return t, nil
	case map[string]interface{}:
		return &jsonObject{
			value: t,
		}, nil
	default:
		return nil, fmt.Errorf("data is %T, not string/[]byte/JSONObject/map[string]interface{}", v)
	}
}

func (obj *jsonObject) GetJSONArray(key string) (JSONArray, error) {
	if !obj.ContainsKey(key) {
		return nil, fmt.Errorf("key %s not exsits", key)
	}
	v := obj.value[key]
	switch t := v.(type) {
	case string:
		return NewJSONArrayFrom([]byte(t))
	case []byte:
		return NewJSONArrayFrom(t)
	case JSONArray:
		return t, nil
	case []interface{}:
		return &jsonArray{
			value: t,
		}, nil
	default:
		return nil, fmt.Errorf("data is %T, not string/[]byte/JSONArray/[]interface{}", v)
	}
}

func (obj *jsonObject) Put(key string, value interface{}) error {
	obj.value[key] = value
	return nil
}

func (obj *jsonObject) GetValue() (map[string]interface{}, error) {
	return obj.value, nil
}

func (obj *jsonObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(obj.value)
}

func (obj *jsonObject) Scan(dest interface{}) error {
	v, err := obj.MarshalJSON()
	if err != nil {
		return err
	}
	return json.Unmarshal(v, dest)
}
