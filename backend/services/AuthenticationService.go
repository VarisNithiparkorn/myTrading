package services

import (

	dto "github.com/VarisNithiparkorn/cryptoGraph/DTO"
	"github.com/VarisNithiparkorn/cryptoGraph/entities"
	errorsHandle "github.com/VarisNithiparkorn/cryptoGraph/errors"
	"github.com/VarisNithiparkorn/cryptoGraph/repositories"
	"github.com/VarisNithiparkorn/cryptoGraph/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(account *dto.Account) (t *string,e error){
	existedAccount := repositories.FindAccountByUsername(account.Username)
	if existedAccount == nil{
		return nil,errorsHandle.CreateBadRequest(400,"account doesn't exist")
	}
	if !existedAccount.IsActive {
		return nil,errorsHandle.CreateConflictErr(403,"email doesn't verify")
	}
	equal := utils.Compare(account.Password,existedAccount.Password)
	if !equal {
		return nil, errorsHandle.CreateBadRequest(400,"invalid email or password")
	}
	fields:=make(utils.AcceessClaimsFields)
	fields["userId"] = existedAccount.ID.String()
	fields["username"] = existedAccount.Username
	fields["role"] = existedAccount.Role
	token,err := utils.GenerateToken(fields,"accesstoken")
	
	if err!=nil {
		return nil,err
	}
	return &token,nil
}

func CreateAccount(accountDTO dto.Account) error {
	existedAccount := repositories.FindAccountByUsername(accountDTO.Username)
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
	fields := make(utils.EmailClaimsFields)
	fields["userId"] = id
	fields["username"]= account.Username
	token,err := utils.GenerateToken(fields,"emailtoken")
	if (err != nil){
		return errorsHandle.CreateInternalServerErr(500,"Error asserting inserted ID to ObjectID")
	}
	utils.SendHtmlMail([]string{account.Username},"register","template/emailVerificationForm.html",dto.Response{Username: account.Username, Url: "http://localhost:3000/verify-email?token="+ token})
	return nil
}

func VerifyEmail(token string) error{
	claims,err := utils.VerifyToken(token,"emailtoken")
	if(err != nil){
		return errorsHandle.CreateUnAuthorizedErr(401,"Invalid token")
	}
	isExpire := utils.IsExpired(claims)
	if(isExpire){
		return	errorsHandle.CreateUnAuthorizedErr(401,"Token expired")
	}
	account := repositories.FindAccountByUsername(claims.GetUsername())
	account.IsActive = true
	repositories.UpdateAccount(account)
	return nil
}
