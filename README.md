# FASTJSON

`fastjson`是java版本的fastjson库的api做一个翻译，方便习惯java的人操作json数据

- 主要适用场景：层级和字段都不能确定的json
- 这个库并不实现高性能json解析，依赖标准库json
- 这个库并没有100%实现java版的api

## 安装

```bash
go get -u github.com/qauzy/fastjson
```

## Usage

- 解析一段json字符串

```go
import "github.com/qauzy/fastjson"
// json object
jsonObjByte := []byte("{\"beijing\":0.013,\"shanghai\":0.0113,\"guangzhou\":0.0104,\"shenzhen\":0.0117,\"hangzhou\":0.0113,\"chengdu\":0.0114}")
jsonObj, err := NewJSONObject(jsonObjByte)

// json array
jsonArrByte := []byte("[\"beijing\",\"shanghai\",\"guangzhou\",\"shenzhen\",\"hangzhou\",\"chengdu\"]")
jsonArr, err := NewJSONArray(jsonArrByte)
```

- 检测

```go
// 零值检测
isZero := jsonObj.IsZero()
// 包含键检测
isExist := jsonObj.ContainsKey("beijing")
// 包含索引检测
inRange := jsonArr.ContainsIndex(1)
```

- 获取数值
  
``` go
// int
_, err := jsonObj.GetInt("beijing")
// int32
_, err := jsonObj.GetInt32("beijing")
// int64
_, err := jsonObj.GetInt64("beijing")
// float32
_, err := jsonObj.GetFloat32("shanghai")
// float64
_, err := jsonObj.GetFloat64("shanghai")
```

- 获取字符串

```go
// string
_, err := jsonObj.GetString("beijing")
```

- 获取bool值
  
```go
// bool
_, err := jsonObj.GetBool("beijing")
```

- 获取嵌套

```go
// JSONObject
_, err := jsonObj.GetJSONObject("beijing")
// JSONArray
_, err := jsonObj.GetJSONArray("beijing")
```

- 更新值
  
```go
// JSONObject
err := jsonObj.Put("wuhan", 0.321)
// JSONArray 追加
err := jsonArr.Put("wuhan")
```

- 转json字符串

```go
_, err := json.Marshal(jsonObj)
```

- 转结构体
  
```go
type agerange struct {
    R1 float32 `json:"beijing"`
    R2 float32 `json:"shanghai"`
}
s := new(agerange)
err := jsonObj.Scan(s)
```
