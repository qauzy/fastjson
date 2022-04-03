package fastjson

import (
	"fmt"
	"strconv"
)

func (arr *jsonArray) Size() int {
	return len(arr.value)
}

func (arr *jsonArray) IsZero() bool {
	return arr.Size() == 0
}

func (arr *jsonArray) ContainsIndex(index int) bool {
	return arr.Size() > index
}

func (arr *jsonArray) GetInt(index int) (int, error) {
	if !arr.ContainsIndex(index) {
		return 0, fmt.Errorf("index %d out of range", index)
	}
	v := arr.value[index]
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

func (arr *jsonArray) GetInt32(index int) (int32, error) {
	if !arr.ContainsIndex(index) {
		return 0, fmt.Errorf("index %d out of range", index)
	}
	v := arr.value[index]
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

func (arr *jsonArray) GetInt64(index int) (int64, error) {
	if !arr.ContainsIndex(index) {
		return 0, fmt.Errorf("index %d out of range", index)
	}
	v := arr.value[index]
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

func (arr *jsonArray) GetFloat32(index int) (float32, error) {
	if !arr.ContainsIndex(index) {
		return 0, fmt.Errorf("index %d out of range", index)
	}
	v := arr.value[index]
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

func (arr *jsonArray) GetFloat64(index int) (float64, error) {
	if !arr.ContainsIndex(index) {
		return 0, fmt.Errorf("index %d out of range", index)
	}
	v := arr.value[index]
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

func (arr *jsonArray) GetString(index int) (string, error) {
	if !arr.ContainsIndex(index) {
		return "", fmt.Errorf("index %d out of range", index)
	}
	v := arr.value[index]
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

func (arr *jsonArray) GetBool(index int) (bool, error) {
	if !arr.ContainsIndex(index) {
		return false, fmt.Errorf("index %d out of range", index)
	}
	v := arr.value[index]
	t, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf("data is %T, not bool", v)
	}
	return t, nil
}

func (arr *jsonArray) GetJSONObject(index int) (JSONObject, error) {
	if !arr.ContainsIndex(index) {
		return nil, fmt.Errorf("index %d out of range", index)
	}
	v := arr.value[index]
	switch t := v.(type) {
	case string:
		return NewJSONObject([]byte(t))
	case []byte:
		return NewJSONObject(t)
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

func (arr *jsonArray) GetJSONArray(index int) (JSONArray, error) {
	if !arr.ContainsIndex(index) {
		return nil, fmt.Errorf("index %d out of range", index)
	}
	v := arr.value[index]
	switch t := v.(type) {
	case string:
		return NewJSONArray([]byte(t))
	case []byte:
		return NewJSONArray(t)
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

func (arr *jsonArray) Put(value interface{}) error {
	arr.value = append(arr.value, value)
	return nil
}

func (arr *jsonArray) GetValue() ([]interface{}, error) {
	return arr.value, nil
}

func (arr *jsonArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(arr.value)
}

func (arr *jsonArray) Scan(dest []interface{}) error {
	v, err := arr.MarshalJSON()
	if err != nil {
		return err
	}
	return json.Unmarshal(v, &dest)
}
