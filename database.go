package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbClient *mongo.Client
	dbCtx    context.Context
	dbCursor *mongo.Database
	dbCol    *mongo.Collection
	dbCancel context.CancelFunc
)

func initDB(ctx context.Context) {
	dbCtx = ctx

	dbString := fmt.Sprintf("mongodb://%s:%s", mongoHostname, mongoPort)
	fmt.Println(dbString, "dbString")
	clientOption := options.Client().ApplyURI(dbString)
	client, err := mongo.Connect(dbCtx, clientOption)
	if err != nil {
		log.Fatal(err, "++++++++=")
	}
	dbClient = client
	dbCursor = dbClient.Database("telegram")
	fmt.Println(dbCursor, "DBCURSOR")
	dbCol = dbCursor.Collection("requests")

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
}
