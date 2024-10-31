package utils

import (
	"math/rand"
	"time"
)

func SliceIsMemberIndex(s []int32, e int32) (int32, bool) {
	for index, i := range s {
		if e == i {
			return int32(index), true
		}
	}
	return -1, false
}

//SliceIsMember slice中是否包含该元素
func SliceIsMember(s []int32, e int32) bool {
	for _, i := range s {
		if e == i {
			return true
		}
	}
	return false
}

// SliceIndexMember 包含哪一个位置的元素
func SliceIndexMember(s []int32, e int32) (int32, bool) {
	for index, i := range s {
		if e == i {
			return int32(index), true
		}
	}
	return -1, false
}

//SliceIsMember64 slice中是否包含该元素
func SliceIsMember64(s []int64, e int64) bool {
	for _, i := range s {
		if e == i {
			return true
		}
	}
	return false
}

//SliceIsContains s1是否完全包含在s2
func SliceIsContains(s1 []int32, s2 []int32) bool {
	for _, item := range s1 {
		found := false
		for _, value := range s2 {
			if item == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

//SliceCopyAppend slice拷贝追加
func SliceCopyAppend(base, add []int32) []int32 {
	res := make([]int32, 0, len(base)+len(add))
	res = append(res, base...)
	res = append(res, add...)
	return res
}

//SliceGet 获取切片元素
func SliceGet(index int32, base []int32) (int32, bool) {
	if index < 0 {
		return 0, false
	}
	if int32(len(base)) < index+1 {
		return 0, false
	}
	return base[index], true
}

//SliceGetSafe 安全获取切片元素,如果没找到就获取最后一个
func SliceGetSafe(index int32, base []int32) int32 {
	if index < 0 {
		return base[0]
	}
	if int32(len(base)) < index+1 {
		return base[len(base)-1]
	}
	return base[index]
}

//SliceIndexRangeScore 目标位于哪个索引 [1,20,30] >=20&<30 index=1
func SliceIndexRangeScore(target int32, slice []int32) int32 {
	if target <= slice[0] {
		//比最小的小
		return 0
	}
	if target >= slice[len(slice)-1] {
		//比最大的大
		return int32(len(slice) - 1)
	}
	for index, nums := range slice {
		if target < nums {
			return int32(index) - 1
		}
	}
	return int32(len(slice) - 1)
}

//SliceRem 删除切片元素
func SliceRem(arr []int32, elem int32) []int32 {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			arr = append(arr[:i], arr[i+1:]...)
			return arr
		}
	}
	return arr
}

//SliceRem64 删除切片元素
func SliceRem64(arr []int64, elem int64) []int64 {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			arr = append(arr[:i], arr[i+1:]...)
			return arr
		}
	}
	return arr
}

//SliceShuffle 打乱数组
func SliceShuffle(base []int32) []int32 {
	res := make([]int32, len(base))
	rand.Seed(time.Now().UnixNano())
	for i, index := range rand.Perm(len(base)) {
		res[i] = base[index]
	}
	return res
}

//SliceShuffleByStr 打乱数组
func SliceShuffleByStr(base []string) []string {
	res := make([]string, len(base))
	rand.Seed(time.Now().UnixNano())
	for i, index := range rand.Perm(len(base)) {
		res[i] = base[index]
	}
	return res
}

//SliceMax 获取最大值
func SliceMax(slice []int32) int32 {
	max := slice[0]
	for _, num := range slice {
		if num > max {
			max = num
		}
	}
	return max
}

func SliceRemDuplicates(slice []int32) []int32 {
	encountered := map[int32]bool{}
	result := []int32{}
	for v := range slice {
		if encountered[slice[v]] == true {
			continue
		} else {
			encountered[slice[v]] = true
			result = append(result, slice[v])
		}
	}
	return result
}

//SliceTotal 求和
func SliceTotal(slice []int32) int32 {
	sum := int32(0)
	for _, num := range slice {
		sum += num
	}
	return sum
}

// SliceIsRepeat 是否有重复slice
func SliceIsRepeat(data []int32) bool {
	check := make(map[int32]int32, len(data))
	for _, datum := range data {
		check[datum]++
		if check[datum] > 1 {
			return true
		}
	}
	return false
}

// SliceIsRepeatInt64 是否有重复slice
func SliceIsRepeatInt64(data []int64) bool {
	check := make(map[int64]int32, len(data))
	for _, datum := range data {
		check[datum]++
		if check[datum] > 1 {
			return true
		}
	}
	return false
}
