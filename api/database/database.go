package database

import (
	"context"
	"fmt"
	"os"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
const (
	databaseName string = "AdaptDB"
	locationsCollection string = "locations"
	savedGamesCollection string = "saved_games"
)

func Init() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	
	if err != nil {
		panic(err)
	}
	
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func GetDB() *mongo.Client{
	return client
}

//Retrieve the "locations" collection from the database.
func GetLocationsCollection() *mongo.Collection {
	return client.Database(databaseName).Collection(locationsCollection)
}

//Retrieve the "saved_games" collection from the database.
func GetSavedGamesCollections() *mongo.Collection {
	return client.Database(databaseName).Collection(savedGamesCollection)
}

func DisconnectClient(){
	client.Disconnect(context.TODO())
}