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

func ConnectToDB(addr string) {
	var err error
	DB, err = mongo.Connect(options.Client().ApplyURI("mongodb://" + addr))
	if err != nil {
		panic("failed to connect to db!")
	}
}

func PingDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := DB.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("failed to ping database!")
	}
}
