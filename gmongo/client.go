package gmongo

import (
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	RoleId int
	Name   string
	Age    string
}

func (this_ *User) CollectionName() string {
	return "user"
}

func (this_ *User) PrimaryKey() bson.M {
	return bson.M{"roleId": this_.RoleId}
}

func (this_ *User) Bson() bson.M {
	return bson.M{"name": this_.Name, "age": this_.Age}
}
