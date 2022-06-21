package gredis

import (
	"reflect"
	"testing"
)

func Test_getClient(t *testing.T) {
	url := "test4pubplat@10.0.240.234:6379/1"
	client, _ := getClient(url)
	tests := []struct {
		name    string
		url     string
		want    *Client
		wantErr bool
	}{
		{
			name:    "bpc 内网测试",
			url:     "test4pubplat@10.0.240.234:6379/1",
			want:    client,
			wantErr: false,
		},
		{
			name:    "bpc 内网重复测试",
			url:     "test4pubplat@10.0.240.234:6379/1",
			want:    client,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getClient(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
