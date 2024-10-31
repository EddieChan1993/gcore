package utils

import (
	"math"
	"reflect"
)

//GetModName 获取结构体名
func GetModName(i interface{}) string {
	return reflect.TypeOf(i).String()
}

//Cal64Safe 安全加减运算
//num计算结果
//isOver是否溢出
func Cal64Safe(left, right int64) (num int64, isOver bool) {
	if right > 0 {
		if left > math.MaxInt64-right {
			return 0, true
		}
	} else {
		if left < math.MinInt64-right {
			return 0, true
		}
	}
	return left + right, false
}

//Cal32Safe 安全加减运算
//num计算结果
//isOver是否溢出
func Cal32Safe(left, right int32) (num int32, isOver bool) {
	if right > 0 {
		if left > math.MaxInt32-right {
			return 0, true
		}
	} else {
		if left < math.MinInt32-right {
			return 0, true
		}
	}
	return left + right, false
}

func Max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func Max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func Min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
