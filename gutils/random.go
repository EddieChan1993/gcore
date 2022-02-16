package gutils

import (
	"errors"
	"math/rand"
	"reflect"
	"time"
	"unsafe"
)

var (
	ErrUnkown = errors.New("found a bug")
	ErrLength = errors.New("param length is need to be greater than 0")
	ErrWeight = errors.New("wieght slice element need to be greater than 0")

	stdStrCommon    = "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	stdStrLowercase = "abcdefghijklmnopqrstuvwsyz"
	stdStrUpercase  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	stdStrDigit     = "0123456789"
)

const (
	requestIDLength = 32
)

var ra = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandRequestID 生成一个32字节长的RequestID
func RandRequestID() string {
	s, _ := RandString(requestIDLength)
	return s
}

// RandString 生成固定长度的字符串，包含大小写、数字，要求lenght的长度大于0
func RandString(length int) (string, error) {
	return randstr(&stdStrCommon, length)
}

// RandDigitString 生成纯数字字符串
func RandDigitString(length int) (string, error) {
	return randstr(&stdStrDigit, length)
}

// RandUpercaseString 生成大写字符串
func RandUpercaseString(length int) (string, error) {
	return randstr(&stdStrUpercase, length)
}

// RandLowercaseString 生成小写字符串
func RandLowercaseString(length int) (string, error) {
	return randstr(&stdStrLowercase, length)
}

// RandWeightIndex 给定一个带权重的slice，返回命中结果的下标，要求weight的长度大于0，并且slice内部的元素要求大于0
func RandWeightIndex(weight []int) (index int, err error) {
	if len(weight) == 0 {
		return 0, ErrLength
	}

	total := 0
	for i := 0; i < len(weight); i++ {
		if weight[i] <= 0 {
			return 0, ErrWeight
		}
		total += weight[i]
	}
	t := ra.Intn(total)
	past := 0
	for i := 0; i < len(weight); i++ {
		past += weight[i]
		if t < past {
			return i, nil
		}
	}

	// 这块儿永远不会走到
	return 0, ErrUnkown
}

func bytesToString(b []byte) (s string) {
	_bptr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	_sptr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	_sptr.Data = _bptr.Data
	_sptr.Len = _bptr.Len
	return s
}

func randstr(std *string, length int) (string, error) {
	if length <= 0 {
		return "", ErrLength
	}

	bytes := []byte(*std)
	var result []byte = make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = bytes[ra.Intn(len(bytes))]
	}

	return bytesToString(result), nil
}
