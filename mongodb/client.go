package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Mongo struct {
	ctx *context.Context
	client *mongo.Client
}

type Database *mongo.Database

func MongoClient() (*Mongo, error) {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(opts)
	if err != nil{
		return &Mongo{}, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil{
		return &Mongo{}, err
	}
	mongo_client := Mongo{
		ctx:   &ctx,
		client: client,
	}
	return &mongo_client, nil
}
