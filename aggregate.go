package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func RunAggregation(dbName, collName string) error {
	var err error
	client, _ := getMongoClient()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	aggregationPipeline := bson.A{
		bson.D{{"$match", bson.D{{Key: "seq", Value: bson.D{{"$gt", 1000}}}}}},
		bson.D{
			{"$project",
				bson.D{
					{"_id", 1},
					{"type", bson.D{{"$type", "$_id"}}},
				},
			},
		},
		bson.D{{"$sort", bson.D{{"_id", 1}}}},
	}
	var cursor *mongo.Cursor

	sc := client.Database(dbName).Collection(collName)
	if cursor, err = sc.Aggregate(ctx, aggregationPipeline); err != nil {
		fmt.Errorf("error : %v", err)
		return err
	}
	//if cursor, err = sc.Find(ctx, qf, opts); err != nil {
	//	return err
	//}

	type docID struct {
		ID   interface{} `bson:"_id"`
		Type interface{} `bson:"type"`
	}
	var ids []interface{}

	cnt := int64(0)
	for cursor.Next(ctx) {
		cnt++
		var doc docID
		if err = cursor.Decode(&doc); err != nil {
			cursor.Close(ctx)
			fmt.Errorf("error: %v", err)
		}
		ids = append(ids, doc.ID)
		fmt.Printf("type %v ,value : %v", doc.Type, doc.ID)
		fmt.Println()
		time.Sleep(100)
	}
	return err
}
