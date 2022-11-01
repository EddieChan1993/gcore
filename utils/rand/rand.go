package gutils

import (
	"errors"
	"math/rand"
	"reflect"
	"time"
	"unsafe"
)

var (
	ErrLength = errors.New("param length is need to be greater than 0")

	stdStrCommon      = "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	stdStrLowercase   = "abcdefghijklmnopqrstuvwsyz"
	stdStrUpercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	stdStrDigit       = "0123456789"
	stdStrDigitNoZero = "123456789"
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
	return randStr(&stdStrCommon, length)
}

// RandDigitString 生成纯数字字符串
func RandDigitString(length int) (string, error) {
	return randStr(&stdStrDigit, length)
}

// RandDigitString 生成纯数字非0字符串
func RandDigitStringNoZero(length int) (string, error) {
	return randStr(&stdStrDigitNoZero, length)
}

// RandUpercaseString 生成大写字符串
func RandUpercaseString(length int) (string, error) {
	return randStr(&stdStrUpercase, length)
}

// RandLowercaseString 生成小写字符串
func RandLowercaseString(length int) (string, error) {
	return randStr(&stdStrLowercase, length)
}

func bytesToString(b []byte) (s string) {
	_bptr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	_sptr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	_sptr.Data = _bptr.Data
	_sptr.Len = _bptr.Len
	return s
}

func randStr(std *string, length int) (string, error) {
	if length <= 0 {
		return "", ErrLength
	}

	bytes := []byte(*std)
	var result = make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = bytes[ra.Intn(len(bytes))]
	}

	return bytesToString(result), nil
}

const RandTotal = 1000             //默认权重总和
type RandPoolTyp = map[int32]int32 //随机池类型

//RandInt32 返回一个(0,total]的随机数
func RandInt32(total int32) int32 {
	rand.Seed(time.Now().UnixNano())
	num := rand.Int31n(total + 1)
	if num == 0 {
		return 1
	}
	return num
}

//RandWeight 计算权重
func RandWeight(pool RandPoolTyp) (weightList, poolIds []int32, weightTotal int32) {
	poolIds = make([]int32, 0, len(pool))
	weightList = make([]int32, 0, len(pool))
	weightTotal = int32(0)
	for id, w := range pool {
		poolIds = append(poolIds, id)
		weightTotal += w
		weightList = append(weightList, weightTotal)
	}
	return
}

//RandOneOnceWeight 指定权重随机
func RandOneOnceWeight(weightList, poolIds []int32, weightTotal int32) (poolId int32) {
	randInt := RandInt32(weightTotal)
	index := 0
	for i, w := range weightList {
		index = i
		if w >= randInt {
			break
		}
	}
	return poolIds[index]
}

//RandOne 随机产出一个
//pool 奖池；k-奖品id v-奖品权重
func RandOne(pool RandPoolTyp) (poolId int32) {
	poolIds := make([]int32, 0, len(pool))
	weightList := make([]int32, 0, len(pool))
	weightTotal := int32(0)
	for id, w := range pool {
		poolIds = append(poolIds, id)
		weightTotal += w
		weightList = append(weightList, weightTotal)
	}
	randInt := RandInt32(weightTotal)
	index := 0
	for i, w := range weightList {
		index = i
		if w >= randInt {
			break
		}
	}
	return poolIds[index]
}

//RandMulti 随机产出多个
//pool 奖池；k-奖品id v-奖品权重
func RandMulti(pool RandPoolTyp, randTimes int32) (poolRandIds []int32) {
	if int32(len(pool)) < randTimes {
		//权重和奖池不等
		return nil
	}
	poolIds := make([]int32, 0, len(pool))
	weightList := make([]int32, 0, len(pool))
	weightTotal := int32(0)
	for id, w := range pool {
		poolIds = append(poolIds, id)
		weightTotal += w
		weightList = append(weightList, weightTotal)
	}
	poolRandIds = make([]int32, 0, randTimes)
	for i := int32(0); i < randTimes; i++ {
		randInt := RandInt32(weightTotal)
		index := 0
		for j, w := range weightList {
			index = j
			if w >= randInt {
				break
			}
		}
		poolRandIds = append(poolRandIds, poolIds[index])
	}
	return poolRandIds
}

//RandMultiNoRepeat  随机产出N个不重复奖品
//pool 奖池；k-奖品id v-奖品权重
func RandMultiNoRepeat(pool RandPoolTyp, randTimes int32) (poolRandIds []int32) {
	if int32(len(pool)) < randTimes {
		//权重和奖池不等，奖池少于需要随机的个数
		return nil
	} else if int32(len(pool)) == randTimes {
		poolIds := make([]int32, 0, len(pool))
		for id := range pool {
			poolIds = append(poolIds, id)
		}
		return poolIds
	}
	poolRandIds = make([]int32, 0, randTimes)
	for i := int32(0); i < randTimes; i++ {
		poolIds := make([]int32, 0, len(pool))
		weightList := make([]int32, 0, len(pool))
		weightTotal := int32(0)
		for id, w := range pool {
			poolIds = append(poolIds, id)
			weightTotal += w
			weightList = append(weightList, weightTotal)
		}
		randInt := RandInt32(weightTotal)
		index := 0
		for j, w := range weightList {
			index = j
			if w >= randInt {
				break
			}
		}
		delete(pool, poolIds[index])
		poolRandIds = append(poolRandIds, poolIds[index])
	}
	return poolRandIds
}
