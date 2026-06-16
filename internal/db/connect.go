package db

import (
	"context"
	"fmt"
	"time"

	"github.com/mostafizurrahaman/go_server/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect(config *config.Config) (*mongo.Client, *mongo.Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	// ? Db options:
	dbOptions := options.Client().ApplyURI(config.MongoURI)

	// ? Mongo connect:
	mongoClient, err := mongo.Connect(dbOptions)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to connect mongodb")
	}

	// ? Ping to mongo db:
	if err := mongoClient.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("Failed to ping into db.")
	}

	fmt.Println("Mongodb database connected and pinged successfully.")

	// ? Configure database :

	database := mongoClient.Database(config.MongoDB)

	fmt.Println(database, mongoClient)

	return mongoClient, database, nil

}

func Disconnect(client *mongo.Client) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	return client.Disconnect(ctx)

}
