package dbconfig

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbClient *mongo.Client

//GetDB return mongo connection.
func GetDB() *mongo.Client {
	if dbClient == nil {
		dbClient = connect()
	}

	return dbClient
}

//Connect mongodb
func connect() *mongo.Client {

	clientOptions := options.Client().ApplyURI("Mongo-DB-STRING")
	client, err := mongo.NewClient(clientOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	return client
}
