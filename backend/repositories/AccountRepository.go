package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/VarisNithiparkorn/cryptoGraph/config"
	"github.com/VarisNithiparkorn/cryptoGraph/entities"

	// "github.com/VarisNithiparkorn/cryptoGraph/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
func intitial() *mongo.Collection{
	db,err := config.StartConnection()
	if(err != nil){
		log.Println("connect to collection fail" + err.Error())
		return nil
	}
	return db.Collection("accounts")
}

func InsertOneAccount(newAccountDTO entities.Account) interface{}{
	accCollection := intitial()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	newAccountDTO.CreatedAt = time.Now()
	newAccountDTO.UpdatedAt = time.Now()
	defer cancel()
	result,err := accCollection.InsertOne(ctx,newAccountDTO)
	if(err != nil){
		fmt.Errorf("can't create account %w",err)
	}
	return result.InsertedID
}

func FindAccountByUsername(username string) *entities.Account{
	var account entities.Account
	accCollection := intitial()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"username":username}
	err := accCollection.FindOne(ctx,filter).Decode(&account)
	if(err != nil){
		return nil
	}
	log.Println("username"+ account.Username)
	return &account
}

func UpdateAccount(account *entities.Account){
	accCollection := intitial()
	account.UpdatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"username":account.Username}
	if _, err := accCollection.ReplaceOne(ctx, filter, account); err != nil {
		fmt.Errorf("update account fail") 
	}
	
}


