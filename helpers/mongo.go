package helpers

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartMongoDb() *mongo.Database {
	ctxTimeOut := (10 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeOut)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_CONN")))

	defer cancel()

	if err != nil {
		panic(err)
	}

	database := client.Database(os.Getenv("MONGO_DB"))
	if database == nil {
		panic(fmt.Errorf("Database doesn't exist"))
	}

	return database
}