package cast

// ToBool 转换 interface 到 bool
func ToBool(i interface{}) bool {
	v, _ := ToBoolE(i)
	return v
}

// ToInt64 转换 interface 到 int64
func ToInt64(i interface{}) int64 {
	v, _ := ToInt64E(i)
	return v
}

// ToInt32 转换 interface 到 int32
func ToInt32(i interface{}) int32 {
	v, _ := ToInt32E(i)
	return v
}

// ToInt16 转换 interface 到 int16
func ToInt16(i interface{}) int16 {
	v, _ := ToInt16E(i)
	return v
}

// ToInt8 转换 interface 到 int8
func ToInt8(i interface{}) int8 {
	v, _ := ToInt8E(i)
	return v
}

// ToInt 转换 interface 到 int
func ToInt(i interface{}) int {
	v, _ := ToIntE(i)
	return v
}

// ToUint64 转换 interface 到 uin64
func ToUint64(i interface{}) uint64 {
	v, _ := ToUint64E(i)
	return v
}

// ToUint32 转换 interface 到 uint32
func ToUint32(i interface{}) uint32 {
	v, _ := ToUint32E(i)
	return v
}

// ToUint16 转换 interface 到 uint16
func ToUint16(i interface{}) uint16 {
	v, _ := ToUint16E(i)
	return v
}

// ToUint8 转换 interface 到 uint8
func ToUint8(i interface{}) uint8 {
	v, _ := ToUint8E(i)
	return v
}

// ToUint 转换 interface 到 uint
func ToUint(i interface{}) uint {
	v, _ := ToUintE(i)
	return v
}

// ToFloat64 转换 interface 到 float64
func ToFloat64(i interface{}) float64 {
	v, _ := ToFloat64E(i)
	return v
}

// ToFloat32 转换 interface 到 float32
func ToFloat32(i interface{}) float32 {
	v, _ := ToFloat32E(i)
	return v
}

// ToString 转换 interface 到 string
func ToString(i interface{}) string {
	v, _ := ToStringE(i)
	return v
}
