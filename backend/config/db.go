package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


func startConnection()(*mongo.Client, error) {
	var conn = LoadEnv()
	var ctx,  cancel = context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	client, error := mongo.Connect(ctx,options.Client().ApplyURI(conn.URI))
	if error != nil{
		log.Println("can't connect with this URI")
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {

		if disconnectErr := client.Disconnect(context.Background()); disconnectErr != nil {
			log.Printf("Error during disconnect after ping failure: %v", disconnectErr)
		}
	}

	fmt.Println("Successfully connected and pinged MongoDB!")
	return client, nil
}