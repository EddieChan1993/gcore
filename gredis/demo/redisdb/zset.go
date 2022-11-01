//ZAdd 向有序集合添加一个或多个成员，或者更新已存在成员的分数
func ZAdd(key string, members ...redis.Z) (int64, error) {
	res, err := gredis.ZAdd(static.StaticRedisUrl(), key, members...)
	if err != nil {
		return 0, err
	}
	resData, err := res.Result()
	if err != nil {
		return 0, err
	}
	return resData, nil
}

//ZRevRank 返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序
func ZRevRank(key string, member string) (int64, error) {
	res, err := gredis.ZRevRank(static.StaticRedisUrl(), key, member)
	if err != nil {
		return 0, err
	}
	resData, err := res.Result()
	if err != nil {
		return 0, err
	}
	return resData + 1, nil
}

//ZRevRangeWithScores 返回有序集中指定分数区间内的成员，分数从高到低排序
func ZRevRangeWithScores(key string, start, stop int32) ([]redis.Z, error) {
	res, err := gredis.ZRevRangeWithScores(static.StaticRedisUrl(), key, int64(start-1), int64(stop-1))
	if err != nil {
		return nil, err
	}
	resData, err := res.Result()
	if err != nil {
		return nil, err
	}
	return resData, nil
}

//ZScore 返回有序集中，成员的分数值
func ZScore(key string, member string) (int64, error) {
	res, err := gredis.ZScore(static.StaticRedisUrl(), key, member)
	if err != nil {
		return 0, err
	}
	resData, err := res.Result()
	if err != nil {
		return 0, err
	}
	return int64(resData), nil
}

//ZRevRange 返回有序集中指定区间内的成员，通过索引，分数从高到底
func ZRevRange(key string, start, stop int32) ([]string, error) {
	res, err := gredis.ZRevRange(static.StaticRedisUrl(), key, int64(start-1), int64(stop-1))
	if err != nil {
		return nil, err
	}
	resData, err := res.Result()
	if err != nil {
		return nil, err
	}
	return resData, nil
}

//ZCount 计算在有序集合中指定区间分数的成员数
func ZCount(key string, min, max string) (*redis.IntCmd, error) {
	return gredis.ZCount(static.StaticRedisUrl(), key, min, max)
}

//Del 计算在有序集合中指定区间分数的成员及其分数
func Del(key string) (int64, error) {
	res, err := gredis.Del(static.StaticRedisUrl(), key)
	if err != nil {
		return 0, err
	}
	resData, err := res.Result()
	if err != nil {
		return 0, err
	}
	return resData, nil
}
