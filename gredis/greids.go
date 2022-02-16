package gredis

import (
	"github.com/go-redis/redis"
	"time"
)

func Get(url string, key string) (*redis.StringCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.Get(key), nil
}

func Set(url string, key string, value interface{}, expiration time.Duration) (*redis.StatusCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.Set(key, value, expiration), nil
}

// ------------------------------------------------------

func HDel(url string, key string, fields ...string) (*redis.IntCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HDel(key, fields...), nil
}

func HExists(url string, key string, field string) (*redis.BoolCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HExists(key, field), nil
}

func HGet(url string, key, field string) (*redis.StringCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HGet(key, field), nil
}

func HGetAll(url string, key string) (*redis.StringStringMapCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HGetAll(key), nil
}

func HIncrBy(url string, key, field string, incr int64) (*redis.IntCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HIncrBy(key, field, incr), nil
}

func HIncrByFloat(url string, key, field string, incr float64) (*redis.FloatCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HIncrByFloat(key, field, incr), nil
}

func HKeys(url string, key string) (*redis.StringSliceCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HKeys(key), nil
}

func HLen(url string, key string) (*redis.IntCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HLen(key), nil
}

func HMGet(url string, key string, fields ...string) (*redis.SliceCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HMGet(key, fields...), nil
}

func HMSet(url string, key string, fields map[string]interface{}) (*redis.StatusCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HMSet(key, fields), nil
}

func HSet(url string, key, field string, value interface{}) (*redis.BoolCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HSet(key, field, value), nil
}

func HSetNX(url string, key, field string, value interface{}) (*redis.BoolCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HSetNX(key, field, value), nil
}

func HVals(url string, key string) (*redis.StringSliceCmd, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.HVals(key), nil
}
