package gredis

import (
	"github.com/EddieChan1993/gcore/utils/cast"
	"strings"
	"sync"

	"github.com/go-redis/redis"
)

type Client struct {
	*redis.Client
}

var lock = sync.RWMutex{}
var clients = map[string]*Client{}

func getClient(url string) (*Client, error) {
	lock.RLock()
	client, ok := clients[url]
	lock.RUnlock()
	if ok {
		return client, nil
	}

	client, err := newClient(url)
	if err != nil {
		return nil, err
	}
	lock.Lock()
	clients[url] = client
	lock.Unlock()
	return client, nil
}

func newClient(url string) (*Client, error) {
	password, host, db := parseURL(url)
	redisOption := &redis.Options{
		Addr:     host,
		DB:       db,
		Password: password,
	}
	c := redis.NewClient(redisOption)
	err := c.Ping().Err()
	if err != nil {
		return nil, err
	}

	return &Client{c}, nil
}

func parseURL(url string) (string, string, int) {
	noHead := parseHeader(url)
	pass, noPass := parsePassword(noHead)
	host, db := parseHost(noPass)
	return pass, host, db
}

func parseHeader(url string) string {
	return strings.TrimPrefix(url, "redis://")
}

func parsePassword(url string) (password, noPassword string) {
	idx := strings.LastIndex(url, "@")
	if idx == -1 {
		return "", url
	}
	return url[:idx], url[idx+1:]
}

func parseHost(url string) (host string, db int) {
	idx := strings.LastIndex(url, "/")
	if idx == -1 {
		return url, 0
	}
	host = url[:idx]
	dbStr := url[idx+1:]
	db = cast.ToInt(dbStr)
	return
}
