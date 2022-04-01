package gmongo

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestMongodb(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	InitDb(ctx, WithUrl("mongodb://localhost:27017"), WithDbName("hat"))
	err := GDb.CreateCollect(ctx, "user")
	if err != nil {
		log.Fatal(err)
	}
	res := &User{RoleId: 1233}
	isExtra, errMongo := GDb.FindOne(ctx, res)
	if errMongo != nil {
		log.Fatal(errMongo)
	}
	log.Println(isExtra, res.Age, res.RoleId)
	errMongo = GDb.Save(ctx, res)
	if errMongo != nil {
		log.Fatal(errMongo)
	}
}
