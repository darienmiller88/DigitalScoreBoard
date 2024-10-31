package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
const (
	DatabaseName string = "AdaptDB"
	LocationsCollection string = "locations"
	SavedGamesCollection string = "saved_games"
)

// builder = "paketobuildpacks/builder:base"
//   buildpacks = ["gcr.io/paketo-buildpacks/go"]

func Init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	
	if err != nil {
		panic(err)
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func GetDB() *mongo.Client{
	return client
}

func DisconnectClient(){
	client.Disconnect(context.TODO())
}