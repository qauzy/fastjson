package fastjson

import (
	"reflect"
	"testing"
)

func Test_jsonArray_Size(t *testing.T) {
	type fields struct {
		value []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name:   "0",
			fields: fields{},
			want:   0,
		},
		{
			name: "not 0",
			fields: fields{
				value: []interface{}{
					0,
					1,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := &jsonArray{
				value: tt.fields.value,
			}
			if got := arr.Size(); got != tt.want {
				t.Errorf("jsonArray.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonArray_IsZero(t *testing.T) {
	type fields struct {
		value []interface{}
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
				value: []interface{}{
					0,
					1,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := &jsonArray{
				value: tt.fields.value,
			}
			if got := arr.IsZero(); got != tt.want {
				t.Errorf("jsonArray.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonArray_ContainsIndex(t *testing.T) {
	type fields struct {
		value []interface{}
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "contain",
			fields: fields{
				value: []interface{}{
					0,
					1,
				},
			},
			args: args{
				index: 1,
			},
			want: true,
		},
		{
			name: "not contain",
			fields: fields{
				value: []interface{}{
					0,
					1,
				},
			},
			args: args{
				index: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := &jsonArray{
				value: tt.fields.value,
			}
			if got := arr.ContainsIndex(tt.args.index); got != tt.want {
				t.Errorf("jsonArray.ContainsIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonArray_GetInt(t *testing.T) {
	type fields struct {
		value []interface{}
	}
	type args struct {
		index int
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
				value: []interface{}{
					"0",
					"18",
				},
			},
			args: args{
				index: 1,
			},
			want: 18,
		},
		{
			name: "float",
			fields: fields{
				value: []interface{}{
					"0",
					18.0,
				},
			},
			args: args{
				index: 1,
			},
			want: 18,
		},
		{
			name: "error",
			fields: fields{
				value: []interface{}{
					"0",
					struct{}{},
				},
			},
			args: args{
				index: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := &jsonArray{
				value: tt.fields.value,
			}
			got, err := arr.GetInt(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonArray.GetInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("jsonArray.GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonArray_GetJSONObject(t *testing.T) {
	type fields struct {
		value []interface{}
	}
	type args struct {
		index int
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
				value: []interface{}{
					"0",
					"{\"beijing\":0.013,\"shanghai\":0.0113,\"guangzhou\":0.0104,\"shenzhen\":0.0117,\"hangzhou\":0.0113,\"chengdu\":0.0114}",
				},
			},
			args: args{
				index: 1,
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
		},
		{
			name: "map[string]interface{}",
			fields: fields{
				value: []interface{}{
					"0",
					map[string]interface{}{
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
				index: 1,
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
		},
		{
			name: "jsonobject",
			fields: fields{
				value: []interface{}{
					"0",
					&jsonObject{
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
				index: 1,
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := &jsonArray{
				value: tt.fields.value,
			}
			got, err := arr.GetJSONObject(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonArray.GetJSONObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonArray.GetJSONObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonArray_GetJSONArray(t *testing.T) {
	type fields struct {
		value []interface{}
	}
	type args struct {
		index int
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
				value: []interface{}{
					"0",
					"[\"beijing\",\"shanghai\",\"guangzhou\",\"shenzhen\",\"hangzhou\",\"chengdu\"]",
				},
			},
			args: args{
				index: 1,
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
		},
		{
			name: "[]interface{}",
			fields: fields{
				value: []interface{}{
					"0",
					[]interface{}{
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
				index: 1,
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
		},
		{
			name: "jsonarray",
			fields: fields{
				value: []interface{}{
					"0",
					&jsonArray{
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
				index: 1,
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := &jsonArray{
				value: tt.fields.value,
			}
			got, err := arr.GetJSONArray(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonArray.GetJSONArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonArray.GetJSONArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonArray_Put(t *testing.T) {
	type fields struct {
		value []interface{}
	}
	type args struct {
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
				value: []interface{}{
					"hangzhou",
					"chengdu",
				},
			},
			args: args{
				value: "wuhan",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := &jsonArray{
				value: tt.fields.value,
			}
			if err := arr.Put(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("jsonArray.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
