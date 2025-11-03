package repositories

import (
	"context"
	"log"
	"time"

	"github.com/VarisNithiparkorn/cryptoGraph/backend/config"
	"github.com/VarisNithiparkorn/cryptoGraph/backend/entities"
	"go.mongodb.org/mongo-driver/mongo"
)
func intitial() *mongo.Collection{
	db,err := config.StartConnection()
	if(err != nil){
		log.Println("connect to collection fail" + err.Error())
	}
	return db.Collection("accounts")
}

func InsertOneAccount(newAccountDTO entities.Account){
	accCollection := intitial()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	newAccountDTO.CreatedAt = time.Now()
	newAccountDTO.UpdatedAt = time.Now()
	defer cancel()
	accCollection.InsertOne(ctx,newAccountDTO)
}


