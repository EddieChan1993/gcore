package gredis

import (
	"github.com/gcore/utils/cast"
	"github.com/go-redis/redis"
	"strings"
)

type Client struct {
	c *redis.Client
}

var clients = map[string]*Client{}

func getClient(url string) (*redis.Client, error) {
	if cli, ok := clients[url]; ok {
		return cli.c, nil
	} else {
		if err := newClient(url); err != nil {
			return nil, err
		}
		cli, _ = clients[url]
		return cli.c, nil
	}
}

func newClient(url string) error {
	if clients[url] != nil {
		return nil
	}
	password, host, db := parseURL(url)
	redisOption := &redis.Options{
		Addr:     host,
		DB:       db,
		Password: password,
	}
	client := redis.NewClient(redisOption)
	err := client.Ping().Err()
	if err != nil {
		return err
	}

	clients[url] = &Client{c: client}
	return nil
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
