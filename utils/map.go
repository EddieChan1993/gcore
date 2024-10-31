package utils

//MapMergeSum map合并，值累加
func MapMergeSum(base, add map[int32]int64) map[int32]int64 {
	if len(base) == 0 {
		return add
	}
	for k, v := range add {
		if old, had := base[k]; had {
			base[k] = old + v
		} else {
			base[k] = v
		}
	}
	return base
}

//Map32MergeSum map合并，值累加
func Map32MergeSum(base, add map[int32]int32) map[int32]int32 {
	if len(base) == 0 {
		return add
	}
	for k, v := range add {
		if old, had := base[k]; had {
			base[k] = old + v
		} else {
			base[k] = v
		}
	}
	return base
}

//CopyMapInt64 map复制
func CopyMapInt64(from map[int32]int64) map[int32]int64 {
	res := make(map[int32]int64, len(from))
	for k, v := range from {
		res[k] = v
	}
	return res
}

//CopyMapInt32 map复制
func CopyMapInt32(from map[int32]int32) map[int32]int32 {
	res := make(map[int32]int32, len(from))
	for k, v := range from {
		res[k] = v
	}
	return res
}

func MapMaxKey(m map[int32]bool) int32 {
	maxKey := int32(0)
	for k := range m {
		if k > maxKey {
			maxKey = k
		}
	}
	return maxKey
}

func Map2Slice(m map[int32]int64) []int32 {
	res := make([]int32, 0, len(m)*2)
	for cfgId, num := range m {
		for i := int64(0); i < num; i++ {
			res = append(res, cfgId)
		}
	}
	return res
}

func MapCopy(m map[int32]int64) map[int32]int64 {
	tmp := make(map[int32]int64, len(m))
	for k, v := range m {
		tmp[k] = v
	}
	return tmp
}
