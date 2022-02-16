package gmongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Init(url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
	defer cancel()
	option := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, option)
	if err != nil {
		return fmt.Errorf("connect err:%v", err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("ping err:%v", err)
	}
	clients[url] = &Client{cli: client, dbs: map[string]*db{}}
	return nil
}

func InitMany(urls []string) error {
	for _, url := range urls {
		if err := Init(url); err != nil {
			return err
		}
	}
	return nil
}

func CreateIndexes(url, db, table string, models []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	col, err := getCollection(url, db, table)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return col.Indexes().CreateMany(ctx, models, opts...)
}

func InsertOne(url, db, table string, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	col, err := getCollection(url, db, table)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return col.InsertOne(ctx, document, opts...)
}

func InsertMany(url, db, table string, document []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	col, err := getCollection(url, db, table)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return col.InsertMany(ctx, document, opts...)
}

func DeleteOne(url, db, table string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	col, err := getCollection(url, db, table)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return col.DeleteOne(ctx, filter, opts...)
}

func DeleteMany(url, db, table string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	col, err := getCollection(url, db, table)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return col.DeleteMany(ctx, filter, opts...)
}

func UpdateOne(url, db, table string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	col, err := getCollection(url, db, table)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return col.UpdateOne(ctx, filter, update, opts...)
}

func UpdateMany(url, db, table string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	col, err := getCollection(url, db, table)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return col.UpdateMany(ctx, filter, update, opts...)
}

func Find(url, db, table string, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	col, err := getCollection(url, db, table)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return col.Find(ctx, filter, opts...)
}

func FindOne(url, db, table string, filter interface{}, opts ...*options.FindOneOptions) (*mongo.SingleResult, error) {
	col, err := getCollection(url, db, table)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return col.FindOne(ctx, filter, opts...), nil
}

func FindOneAndUpdate(url, db, table string, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) (*mongo.SingleResult, error) {
	col, err := getCollection(url, db, table)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return col.FindOneAndUpdate(ctx, filter, update, opts...), nil
}

func CountDocuments(url, db, table string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	col, err := getCollection(url, db, table)
	if err != nil {
		return 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return col.CountDocuments(ctx, filter, opts...)
}
