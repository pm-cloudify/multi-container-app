package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var DB *mongo.Client = nil
var Collection *mongo.Collection = nil
var DBName string
var CollectionName string

func ConnectToDB(addr string) {
	var err error
	DB, err = mongo.Connect(options.Client().ApplyURI("mongodb://" + addr))
	if err != nil {
		panic("failed to connect to db!")
	}

}

func DisconnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := DB.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func PingDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := DB.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("failed to ping database!")
	}
}

func GetCollection(db_name, collection_name string) *mongo.Collection {
	if db_name == DBName && CollectionName == collection_name && Collection != nil {
		return Collection
	}
	DBName = db_name
	CollectionName = collection_name
	Collection = DB.Database(db_name).Collection(collection_name)
	return Collection
}
