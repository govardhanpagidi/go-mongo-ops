package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"time"
)

func PopulateData(dbName, collName string, noOfDocs int64) error {

	client, err := getMongoClient()
	if err != nil {
		return err
	}
	var collection *mongo.Collection
	var ctx = context.Background()
	client.Database(dbName).Collection(collName).Drop(ctx)

	for i := int64(0); i < noOfDocs; i++ {

		var id interface{}
		if i%4 == 0 {
			id = i
		} else if i%3 == 0 {
			id = bson.M{"_id": i, "fld0": RandomString(), "rand": int32(1998)}
		} else if i%2 == 0 {
			id = primitive.NewObjectID()
		} else {
			id = i
		}
		var doc = bson.M{"_id": id, "fld0": "Atlanta", "num": RandomInt(), "seq": i}
		collection = client.Database(dbName).Collection(collName)
		if _, err := collection.InsertOne(ctx, doc); err != nil {
			fmt.Errorf("error:%v", err)
		}
	}

	return err
}

func RandomString() string {
	rand.Seed(time.Now().UnixNano())
	charset := "this is a random string generation used for data simulation."
	c := charset[rand.Intn(len(charset))]
	return string(c)
}

func RandomInt() int {
	return rand.Intn(10000000)
}
