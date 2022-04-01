package gmongo

import (
	"context"
	"errors"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"unsafe"
)

var GDb *mongoDb

const uri = "mongodb://localhost:27017"
const dbName = "demo"

var (
	ErrNoPrimary = errors.New("no primary")
)

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

// InitDb
//consul.ConfigCenter.RegisterConfig("mongo", "mongo")
//consul.ConfigCenter.RegisterConfig("db", "db")
//conf := &qmgo.Config{
//	Uri:      consul.ConfigCenter.GetConfig("mongo").GetString("url"),
//	Database: consul.ConfigCenter.GetConfig("db").GetString(os.Getenv(env.SERVER_ID)),
//}
//gDb = client.Database(consul.ConfigCenter.GetConfig("db").GetString(os.Getenv(env.SERVER_ID)))
func InitDb(ctx context.Context, ops ...Option) {
	if ctx == nil {
		ctx = context.TODO()
	}

	var opObj = &options{
		dbName: dbName,
		url:    uri,
	}
	for _, op := range ops {
		op.apply(opObj)
	}
	conf := &qmgo.Config{
		Uri:      opObj.url,
		Database: opObj.dbName,
	}
	client, err := qmgo.NewClient(ctx, conf)
	if err != nil {
		panic(err)
	}
	GDb = (*mongoDb)(unsafe.Pointer(client.Database(dbName)))
}

func (this_ *mongoDb) GetDb() *qmgo.Database {
	return (*qmgo.Database)(unsafe.Pointer(GDb))
}

//FindOne 查找单条数据
func (this_ *mongoDb) FindOne(ctx context.Context, v IDao) (exist bool, err error) {
	p := v.PrimaryKey()
	if len(p) == 0 {
		return false, ErrNoPrimary
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

//CreateCollect 创建集合
func (this_ *mongoDb) CreateCollect(ctx context.Context, name string) error {
	return this_.GetDb().CreateCollection(ctx, name)
}

//Save 存储集合
func (this_ *mongoDb) Save(ctx context.Context, v IDao) error {
	return this_.GetDb().Collection(v.CollectionName()).UpdateOne(ctx, v.PrimaryKey(), bson.M{"$set": v.Bson()})
}
