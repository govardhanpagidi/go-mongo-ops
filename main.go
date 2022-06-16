package main

import (
	"fmt"
)

var dbName = "test1"
var collName = "coll1"
var connString = "mongodb://localhost:27017"

func main() {
	fmt.Println(dbName, collName)

	err := PopulateData(dbName, collName, 1000)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
}
