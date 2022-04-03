package fastjson

import (
	"reflect"
	"testing"
)

func Test_jsonObject_IsZero(t *testing.T) {
	type fields struct {
		value map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "is zero",
			fields: fields{},
			want:   true,
		},
		{
			name: "not zero",
			fields: fields{
				value: map[string]interface{}{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &jsonObject{
				value: tt.fields.value,
			}
			if got := obj.IsZero(); got != tt.want {
				t.Errorf("jsonObject.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonObject_ContainsKey(t *testing.T) {
	type fields struct {
		value map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "key exist",
			fields: fields{
				value: map[string]interface{}{
					"lang": "go",
				},
			},
			args: args{
				key: "lang",
			},
			want: true,
		},
		{
			name: "key not exist",
			fields: fields{
				value: map[string]interface{}{},
			},
			args: args{
				key: "lang",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &jsonObject{
				value: tt.fields.value,
			}
			if got := obj.ContainsKey(tt.args.key); got != tt.want {
				t.Errorf("jsonObject.ContainsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonObject_GetInt(t *testing.T) {
	type fields struct {
		value map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "string",
			fields: fields{
				value: map[string]interface{}{
					"age": "23",
				},
			},
			args: args{
				key: "age",
			},
			want:    23,
			wantErr: false,
		},
		{
			name: "float",
			fields: fields{
				value: map[string]interface{}{
					"height": 23.8,
				},
			},
			args: args{
				key: "height",
			},
			want:    23,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				value: map[string]interface{}{
					"height": struct{}{},
				},
			},
			args: args{
				key: "height",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &jsonObject{
				value: tt.fields.value,
			}
			got, err := obj.GetInt(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonObject.GetInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("jsonObject.GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonObject_GetJSONObject(t *testing.T) {
	type fields struct {
		value map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    JSONObject
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "string",
			fields: fields{
				value: map[string]interface{}{
					"height": "{\"beijing\":0.013,\"shanghai\":0.0113,\"guangzhou\":0.0104,\"shenzhen\":0.0117,\"hangzhou\":0.0113,\"chengdu\":0.0114}",
				},
			},
			args: args{
				key: "height",
			},
			want: &jsonObject{
				value: map[string]interface{}{
					"beijing":   0.013,
					"shanghai":  0.0113,
					"guangzhou": 0.0104,
					"shenzhen":  0.0117,
					"hangzhou":  0.0113,
					"chengdu":   0.0114,
				},
			},
			wantErr: false,
		},
		{
			name: "map[string]interface{}",
			fields: fields{
				value: map[string]interface{}{
					"height": map[string]interface{}{
						"beijing":   0.013,
						"shanghai":  0.0113,
						"guangzhou": 0.0104,
						"shenzhen":  0.0117,
						"hangzhou":  0.0113,
						"chengdu":   0.0114,
					},
				},
			},
			args: args{
				key: "height",
			},
			want: &jsonObject{
				value: map[string]interface{}{
					"beijing":   0.013,
					"shanghai":  0.0113,
					"guangzhou": 0.0104,
					"shenzhen":  0.0117,
					"hangzhou":  0.0113,
					"chengdu":   0.0114,
				},
			},
			wantErr: false,
		},
		{
			name: "jsonobject",
			fields: fields{
				value: map[string]interface{}{
					"height": &jsonObject{
						value: map[string]interface{}{
							"beijing":   0.013,
							"shanghai":  0.0113,
							"guangzhou": 0.0104,
							"shenzhen":  0.0117,
							"hangzhou":  0.0113,
							"chengdu":   0.0114,
						},
					},
				},
			},
			args: args{
				key: "height",
			},
			want: &jsonObject{
				value: map[string]interface{}{
					"beijing":   0.013,
					"shanghai":  0.0113,
					"guangzhou": 0.0104,
					"shenzhen":  0.0117,
					"hangzhou":  0.0113,
					"chengdu":   0.0114,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &jsonObject{
				value: tt.fields.value,
			}
			got, err := obj.GetJSONObject(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonObject.GetJSONObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonObject.GetJSONObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonObject_GetJSONArray(t *testing.T) {
	type fields struct {
		value map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    JSONArray
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "string",
			fields: fields{
				value: map[string]interface{}{
					"height": "[\"beijing\",\"shanghai\",\"guangzhou\",\"shenzhen\",\"hangzhou\",\"chengdu\"]",
				},
			},
			args: args{
				key: "height",
			},
			want: &jsonArray{
				value: []interface{}{
					"beijing",
					"shanghai",
					"guangzhou",
					"shenzhen",
					"hangzhou",
					"chengdu",
				},
			},
			wantErr: false,
		},
		{
			name: "[]interface{}",
			fields: fields{
				value: map[string]interface{}{
					"height": []interface{}{
						"beijing",
						"shanghai",
						"guangzhou",
						"shenzhen",
						"hangzhou",
						"chengdu",
					},
				},
			},
			args: args{
				key: "height",
			},
			want: &jsonArray{
				value: []interface{}{
					"beijing",
					"shanghai",
					"guangzhou",
					"shenzhen",
					"hangzhou",
					"chengdu",
				},
			},
			wantErr: false,
		},
		{
			name: "jsonarray",
			fields: fields{
				value: map[string]interface{}{
					"height": &jsonArray{
						value: []interface{}{
							"beijing",
							"shanghai",
							"guangzhou",
							"shenzhen",
							"hangzhou",
							"chengdu",
						},
					},
				},
			},
			args: args{
				key: "height",
			},
			want: &jsonArray{
				value: []interface{}{
					"beijing",
					"shanghai",
					"guangzhou",
					"shenzhen",
					"hangzhou",
					"chengdu",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &jsonObject{
				value: tt.fields.value,
			}
			got, err := obj.GetJSONArray(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonObject.GetJSONArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonObject.GetJSONArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonObject_Put(t *testing.T) {
	type fields struct {
		value map[string]interface{}
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "put data",
			fields: fields{
				value: map[string]interface{}{
					"age": "32",
				},
			},
			args: args{
				key:   "height",
				value: 172,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &jsonObject{
				value: tt.fields.value,
			}
			if err := obj.Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("jsonObject.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
