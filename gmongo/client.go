package gmongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Client struct {
	cli *mongo.Client
	dbs map[string]*db
}

type db struct {
	database    *mongo.Database
	collections map[string]*mongo.Collection
}

func (c *Client) getCollection(database, table string) *mongo.Collection {
	if c.dbs[database] == nil {
		c.dbs[database] = &db{database: c.cli.Database(database), collections: map[string]*mongo.Collection{}}
	}

	if c.dbs[database].collections[table] == nil {
		c.dbs[database].collections[table] = c.dbs[database].database.Collection(table)
	}

	return c.dbs[database].collections[table]
}

func getClient(url string) (*Client, error) {
	if clients[url] == nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
		defer cancel()
		option := options.Client().ApplyURI(url)
		client, err := mongo.Connect(ctx, option)
		if err != nil {
			return nil, fmt.Errorf("connect err:%v", err)
		}
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			return nil, fmt.Errorf("ping err:%v", err)
		}
		clients[url] = &Client{cli: client, dbs: map[string]*db{}}
	}
	return clients[url], nil
}

func getCollection(url, db, table string) (*mongo.Collection, error) {
	cli, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return cli.getCollection(db, table), nil
}

var clients = map[string]*Client{}
