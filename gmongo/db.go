package gmongo

import (
	"context"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var GDb *mongoDb

const uri = "mongodb://localhost:27017"
const dbName = "hatgame"

type mongoDb qmgo.Database

type IMongoDb interface {
	GetDb() *qmgo.Database
	FindOne(ctx context.Context, v IDao) (exist bool, err error)
	Save(ctx context.Context, v IDao) error
}

type IDao interface {
	CollectionName() string //集合名字
	PrimaryKey() bson.M     //主键id
	Bson() bson.M           //非主键字段
}

func InitDb(ctx context.Context) {
	if ctx == nil {
		ctx = context.TODO()
	}
	//consul.ConfigCenter.RegisterConfig("mongo", "mongo")
	//consul.ConfigCenter.RegisterConfig("db", "db")
	//conf := &qmgo.Config{
	//	Uri:      consul.ConfigCenter.GetConfig("mongo").GetString("url"),
	//	Database: consul.ConfigCenter.GetConfig("db").GetString(os.Getenv(env.SERVER_ID)),
	//}
	conf := &qmgo.Config{
		Uri:      uri,
		Database: dbName,
	}
	client, err := qmgo.NewClient(ctx, conf)
	if err != nil {
		log.Panic(err)
	}
	//gDb = client.Database(consul.ConfigCenter.GetConfig("db").GetString(os.Getenv(env.SERVER_ID)))
	GDb = (*mongoDb)(client.Database(dbName))
}

func (this_ *mongoDb) GetDb() *qmgo.Database {
	return (*qmgo.Database)(GDb)
}

func (this_ *mongoDb) FindOne(ctx context.Context, v IDao) (exist bool, err error) {
	p := v.PrimaryKey()
	if len(p) == 0 {
		return false, err
	}
	err = this_.GetDb().Collection(v.CollectionName()).Find(ctx, p).One(v)
	if err != nil && err == mongo.ErrNoDocuments {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this_ *mongoDb) Save(ctx context.Context, v IDao) error {
	return this_.GetDb().Collection(v.CollectionName()).UpdateOne(ctx, v.PrimaryKey(), bson.M{"$set": v.Bson()})
}
