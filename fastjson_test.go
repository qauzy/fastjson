package fastjson

import (
	"reflect"
	"testing"
)

func TestNewJSONObject(t *testing.T) {
	type args struct {
		jsonByte []byte
	}
	tests := []struct {
		name    string
		args    args
		want    JSONObject
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "json",
			args: args{
				jsonByte: []byte("{\"beijing\":0.013,\"shanghai\":0.0113,\"guangzhou\":0.0104,\"shenzhen\":0.0117,\"hangzhou\":0.0113,\"chengdu\":0.0114}"),
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
			got, err := NewJSONObject(tt.args.jsonByte)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJSONObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJSONObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJSONArray(t *testing.T) {
	type args struct {
		jsonByte []byte
	}
	tests := []struct {
		name    string
		args    args
		want    JSONArray
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "json",
			args: args{
				jsonByte: []byte("[\"beijing\",\"shanghai\",\"guangzhou\",\"shenzhen\",\"hangzhou\",\"chengdu\"]"),
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
			got, err := NewJSONArray(tt.args.jsonByte)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJSONArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJSONArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
