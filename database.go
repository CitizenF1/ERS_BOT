package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	tele "gopkg.in/telebot.v3"
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
	clientOption := options.Client().ApplyURI(dbString)
	client, err := mongo.Connect(dbCtx, clientOption)
	if err != nil {
		log.Fatal(err)
	}
	dbClient = client
	dbCursor = dbClient.Database("telegram")
	dbCol = dbCursor.Collection("users")

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
}

func storeUsersIntoDB(recipient tele.User) primitive.ObjectID {
	// Return ObjectId
	StoredUser := StoredUsers{User: &recipient}
	res, err := dbCol.InsertOne(dbCtx, StoredUser)
	if err != nil {
		log.Panic(err)
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id
}

func usersFromStoredDB() []StoredUsers {
	usersCollection := dbClient.Database("telegram").Collection("users")
	query := bson.M{}
	users := []StoredUsers{}
	cur, err := usersCollection.Find(dbCtx, query)
	if err != nil {
		fmt.Println(err)
	}
	defer cur.Close(dbCtx)
	cur.All(dbCtx, &users)
	return users
}
