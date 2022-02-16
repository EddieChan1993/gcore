package cast

import (
	"html/template"
	"testing"
)

func TestToBoolE(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    bool
		wantErr bool
	}{
		{0, false, false},
		{nil, false, false},
		{"0", false, false},
		{"false", false, false},
		{"FALSE", false, false},
		{"False", false, false},
		{"f", false, false},
		{"F", false, false},
		{false, false, false},
		{"1", true, false},
		{"true", true, false},
		{"TRUE", true, false},
		{"True", true, false},
		{"t", true, false},
		{"T", true, false},
		{1, true, false},
		{true, true, false},
		{-1, true, false},
		// errors
		{"test", false, true},
		{testing.T{}, false, true},
	}
	for _, tt := range tests {
		t.Run("TestToBoolE", func(t *testing.T) {
			got, err := ToBoolE(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBoolE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToBoolE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64E(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    int64
		wantErr bool
	}{
		{int(2), 2, false},
		{int8(2), 2, false},
		{int16(2), 2, false},
		{int32(2), 2, false},
		{int64(2), 2, false},
		{int(-2), -2, false},
		{int8(-2), -2, false},
		{int16(-2), -2, false},
		{int32(-2), -2, false},
		{int64(-2), -2, false},
		{uint(2), 2, false},
		{uint8(2), 2, false},
		{uint16(2), 2, false},
		{uint32(2), 2, false},
		{uint64(2), 2, false},
		{float32(2.22), 2, false},
		{float64(2.22), 2, false},
		{float32(-2.22), -2, false},
		{float64(-2.22), -2, false},
		{true, 1, false},
		{false, 0, false},
		{"2", 2, false},
		{"-2", -2, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("TestToInt64E", func(t *testing.T) {
			got, err := ToInt64E(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt64E() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32E(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    int32
		wantErr bool
	}{
		{int(2), 2, false},
		{int8(2), 2, false},
		{int16(2), 2, false},
		{int32(2), 2, false},
		{int64(2), 2, false},
		{int(-2), -2, false},
		{int8(-2), -2, false},
		{int16(-2), -2, false},
		{int32(-2), -2, false},
		{int64(-2), -2, false},
		{uint(2), 2, false},
		{uint8(2), 2, false},
		{uint16(2), 2, false},
		{uint32(2), 2, false},
		{uint64(2), 2, false},
		{float32(2.22), 2, false},
		{float64(2.22), 2, false},
		{float32(-2.22), -2, false},
		{float64(-2.22), -2, false},
		{true, 1, false},
		{false, 0, false},
		{"2", 2, false},
		{"-2", -2, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("TestToInt32E", func(t *testing.T) {
			got, err := ToInt32E(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt32E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt32E() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16E(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    int16
		wantErr bool
	}{
		{int(2), 2, false},
		{int8(2), 2, false},
		{int16(2), 2, false},
		{int32(2), 2, false},
		{int64(2), 2, false},
		{int(-2), -2, false},
		{int8(-2), -2, false},
		{int16(-2), -2, false},
		{int32(-2), -2, false},
		{int64(-2), -2, false},
		{uint(2), 2, false},
		{uint8(2), 2, false},
		{uint16(2), 2, false},
		{uint32(2), 2, false},
		{uint64(2), 2, false},
		{float32(2.22), 2, false},
		{float64(2.22), 2, false},
		{float32(-2.22), -2, false},
		{float64(-2.22), -2, false},
		{true, 1, false},
		{false, 0, false},
		{"2", 2, false},
		{"-2", -2, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("TestToInt16E", func(t *testing.T) {
			got, err := ToInt16E(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt16E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt16E() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8E(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    int8
		wantErr bool
	}{
		{int(2), 2, false},
		{int8(2), 2, false},
		{int16(2), 2, false},
		{int32(2), 2, false},
		{int64(2), 2, false},
		{int(-2), -2, false},
		{int8(-2), -2, false},
		{int16(-2), -2, false},
		{int32(-2), -2, false},
		{int64(-2), -2, false},
		{uint(2), 2, false},
		{uint8(2), 2, false},
		{uint16(2), 2, false},
		{uint32(2), 2, false},
		{uint64(2), 2, false},
		{float32(2.22), 2, false},
		{float64(2.22), 2, false},
		{float32(-2.22), -2, false},
		{float64(-2.22), -2, false},
		{true, 1, false},
		{false, 0, false},
		{"2", 2, false},
		{"-2", -2, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("TestToInt8E", func(t *testing.T) {
			got, err := ToInt8E(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt8E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt8E() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToIntE(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    int
		wantErr bool
	}{
		{int(2), 2, false},
		{int8(2), 2, false},
		{int16(2), 2, false},
		{int32(2), 2, false},
		{int64(2), 2, false},
		{int(-2), -2, false},
		{int8(-2), -2, false},
		{int16(-2), -2, false},
		{int32(-2), -2, false},
		{int64(-2), -2, false},
		{uint(2), 2, false},
		{uint8(2), 2, false},
		{uint16(2), 2, false},
		{uint32(2), 2, false},
		{uint64(2), 2, false},
		{float32(2.22), 2, false},
		{float64(2.22), 2, false},
		{float32(-2.22), -2, false},
		{float64(-2.22), -2, false},
		{true, 1, false},
		{false, 0, false},
		{"2", 2, false},
		{"-2", -2, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("TestToIntE", func(t *testing.T) {
			got, err := ToIntE(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIntE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToIntE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64E(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    uint64
		wantErr bool
	}{
		{int(2), 2, false},
		{int8(2), 2, false},
		{int16(2), 2, false},
		{int32(2), 2, false},
		{int64(2), 2, false},
		{uint(2), 2, false},
		{uint8(2), 2, false},
		{uint16(2), 2, false},
		{uint32(2), 2, false},
		{uint64(2), 2, false},
		{float32(2.22), 2, false},
		{float64(2.22), 2, false},
		{true, 1, false},
		{false, 0, false},
		{"2", 2, false},
		{nil, 0, false},
		// errors
		{int(-2), 0, true},
		{int8(-2), 0, true},
		{int16(-2), 0, true},
		{int32(-2), 0, true},
		{int64(-2), 0, true},
		{float32(-2.22), 0, true},
		{float64(-2.22), 0, true},
		{"-2", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("TestToUint64E", func(t *testing.T) {
			got, err := ToUint64E(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint64E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint64E() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32E(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    uint32
		wantErr bool
	}{
		{int(2), 2, false},
		{int8(2), 2, false},
		{int16(2), 2, false},
		{int32(2), 2, false},
		{int64(2), 2, false},
		{uint(2), 2, false},
		{uint8(2), 2, false},
		{uint16(2), 2, false},
		{uint32(2), 2, false},
		{uint64(2), 2, false},
		{float32(2.22), 2, false},
		{float64(2.22), 2, false},
		{true, 1, false},
		{false, 0, false},
		{"2", 2, false},
		{nil, 0, false},
		// errors
		{int(-2), 0, true},
		{int8(-2), 0, true},
		{int16(-2), 0, true},
		{int32(-2), 0, true},
		{int64(-2), 0, true},
		{float32(-2.22), 0, true},
		{float64(-2.22), 0, true},
		{"-2", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("TestToUint32E", func(t *testing.T) {
			got, err := ToUint32E(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint32E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint32E() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16E(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    uint16
		wantErr bool
	}{
		{int(2), 2, false},
		{int8(2), 2, false},
		{int16(2), 2, false},
		{int32(2), 2, false},
		{int64(2), 2, false},
		{uint(2), 2, false},
		{uint8(2), 2, false},
		{uint16(2), 2, false},
		{uint32(2), 2, false},
		{uint64(2), 2, false},
		{float32(2.22), 2, false},
		{float64(2.22), 2, false},
		{true, 1, false},
		{false, 0, false},
		{"2", 2, false},
		{nil, 0, false},
		// errors
		{int(-2), 0, true},
		{int8(-2), 0, true},
		{int16(-2), 0, true},
		{int32(-2), 0, true},
		{int64(-2), 0, true},
		{float32(-2.22), 0, true},
		{float64(-2.22), 0, true},
		{"-2", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("TestToUint16E", func(t *testing.T) {
			got, err := ToUint16E(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint16E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint16E() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint8E(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    uint8
		wantErr bool
	}{
		{int(2), 2, false},
		{int8(2), 2, false},
		{int16(2), 2, false},
		{int32(2), 2, false},
		{int64(2), 2, false},
		{uint(2), 2, false},
		{uint8(2), 2, false},
		{uint16(2), 2, false},
		{uint32(2), 2, false},
		{uint64(2), 2, false},
		{float32(2.22), 2, false},
		{float64(2.22), 2, false},
		{true, 1, false},
		{false, 0, false},
		{"2", 2, false},
		{nil, 0, false},
		// errors
		{int(-2), 0, true},
		{int8(-2), 0, true},
		{int16(-2), 0, true},
		{int32(-2), 0, true},
		{int64(-2), 0, true},
		{float32(-2.22), 0, true},
		{float64(-2.22), 0, true},
		{"-2", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("TestToUint8E", func(t *testing.T) {
			got, err := ToUint8E(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint8E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint8E() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUintE(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    uint
		wantErr bool
	}{
		{int(2), 2, false},
		{int8(2), 2, false},
		{int16(2), 2, false},
		{int32(2), 2, false},
		{int64(2), 2, false},
		{uint(2), 2, false},
		{uint8(2), 2, false},
		{uint16(2), 2, false},
		{uint32(2), 2, false},
		{uint64(2), 2, false},
		{float32(2.22), 2, false},
		{float64(2.22), 2, false},
		{true, 1, false},
		{false, 0, false},
		{"2", 2, false},
		{nil, 0, false},
		// errors
		{int(-2), 0, true},
		{int8(-2), 0, true},
		{int16(-2), 0, true},
		{int32(-2), 0, true},
		{int64(-2), 0, true},
		{float32(-2.22), 0, true},
		{float64(-2.22), 0, true},
		{"-2", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("TestToUintE", func(t *testing.T) {
			got, err := ToUintE(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUintE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUintE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64E(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    float64
		wantErr bool
	}{
		{int(2), 2, false},
		{int8(2), 2, false},
		{int16(2), 2, false},
		{int32(2), 2, false},
		{int64(2), 2, false},
		{int(-2), -2, false},
		{int8(-2), -2, false},
		{int16(-2), -2, false},
		{int32(-2), -2, false},
		{int64(-2), -2, false},
		{uint(2), 2, false},
		{uint8(2), 2, false},
		{uint16(2), 2, false},
		{uint32(2), 2, false},
		{uint64(2), 2, false},
		{float32(2), 2, false},
		// {float32(2.22), 2.22, true},   // 小数位不一致 2.2200000286102295
		{float64(2.22), 2.22, false},
		{float32(-2), -2, false},
		// {float32(-2.22), -2.22, true}, // 小数位不一致 2.2200000286102295
		{float64(-2.22), -2.22, false},
		{true, 1, false},
		{false, 0, false},
		{"2.22", 2.22, false},
		{"-2.22", -2.22, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("TestToFloat64E", func(t *testing.T) {
			got, err := ToFloat64E(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat64E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToFloat64E() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat32E(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    float32
		wantErr bool
	}{
		{int(2), 2, false},
		{int8(2), 2, false},
		{int16(2), 2, false},
		{int32(2), 2, false},
		{int64(2), 2, false},
		{int(-2), -2, false},
		{int8(-2), -2, false},
		{int16(-2), -2, false},
		{int32(-2), -2, false},
		{int64(-2), -2, false},
		{uint(2), 2, false},
		{uint8(2), 2, false},
		{uint16(2), 2, false},
		{uint32(2), 2, false},
		{uint64(2), 2, false},
		{float32(2.22), 2.22, false},
		{float64(2.22), 2.22, false},
		{float32(-2.22), -2.22, false},
		{float64(-2.22), -2.22, false},
		{true, 1, false},
		{false, 0, false},
		{"2.22", 2.22, false},
		{"-2.22", -2.22, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("TestToFloat32E", func(t *testing.T) {
			got, err := ToFloat32E(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat32E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToFloat32E() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStringE(t *testing.T) {
	tests := []struct {
		args    interface{}
		want    string
		wantErr bool
	}{
		{int(2), "2", false},
		{int8(2), "2", false},
		{int16(2), "2", false},
		{int32(2), "2", false},
		{int64(2), "2", false},
		{int(-2), "-2", false},
		{int8(-2), "-2", false},
		{int16(-2), "-2", false},
		{int32(-2), "-2", false},
		{int64(-2), "-2", false},
		{uint(2), "2", false},
		{uint8(2), "2", false},
		{uint16(2), "2", false},
		{uint32(2), "2", false},
		{uint64(2), "2", false},
		{float32(2.22), "2.22", false},
		{float64(2.22), "2.22", false},
		{float32(-2.22), "-2.22", false},
		{float64(-2.22), "-2.22", false},
		{true, "true", false},
		{false, "false", false},
		{nil, "", false},
		{[]byte("abc 123"), "abc 123", false},
		{"abc 123", "abc 123", false},
		{template.HTML("abc 123"), "abc 123", false},
		{template.URL("https://git.dhgames.cn"), "https://git.dhgames.cn", false},
		{template.JS("(1+2)"), "(1+2)", false},
		{template.CSS("abc"), "abc", false},
		{template.HTMLAttr("abc"), "abc", false},
		// errors
		{testing.T{}, "", true},
	}
	for _, tt := range tests {
		t.Run("TestToStringE", func(t *testing.T) {
			got, err := ToStringE(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToStringE() got = %v, want %v", got, tt.want)
			}
		})
	}
}
