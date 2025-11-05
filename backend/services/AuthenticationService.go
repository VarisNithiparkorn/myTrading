package services

import (
	"log"

	dto "github.com/VarisNithiparkorn/cryptoGraph/DTO"
	"github.com/VarisNithiparkorn/cryptoGraph/entities"
	errorsHandle "github.com/VarisNithiparkorn/cryptoGraph/errors"
	"github.com/VarisNithiparkorn/cryptoGraph/repositories"
	"github.com/VarisNithiparkorn/cryptoGraph/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAccount(accountDTO dto.Account) error {
	existedAccount := repositories.FindAccountByUsername(accountDTO.Username)
	log.Println("accouttttttttttttttt", existedAccount)
	if existedAccount != nil {
		err := errorsHandle.CreateConflictErr(409,"username already existed")
		return err
	}
	account := entities.Account{
		Username: accountDTO.Username,
		Password: utils.Hash(accountDTO.Password),
		IsActive:false,
	}
	insertedId := repositories.InsertOneAccount(account)
	objectId,ok := insertedId.(primitive.ObjectID)
	if !ok {
        return errorsHandle.CreateInternalServerErr(500,"Error asserting inserted ID to ObjectID")
    }
	id := objectId.Hex()
	token,err := utils.GenerateAccessToken(id,account.Username)
	if (err != nil){
		return errorsHandle.CreateInternalServerErr(500,"Error asserting inserted ID to ObjectID")
	}
	utils.SendHtmlMail([]string{account.Username},"register","template/emailVerificationForm.html",dto.Response{Username: account.Username, Url: "http://localhost:3000/verify-email?token="+ token})
	return nil
}

func VerifyEmail(token string) error{
	claims,err := utils.VerifyAccessToken(token)
	if(err != nil){
		return errorsHandle.CreateUnAuthorizedErr(401,"Invalid token")
	}
	isExpire := utils.IsExpired(claims)
	if(isExpire){
		return	errorsHandle.CreateUnAuthorizedErr(401,"Token expired")
	}
	account := repositories.FindAccountByUsername(claims.Username)
	account.IsActive = true
	repositories.UpdateAccount(account)
	return nil
}
