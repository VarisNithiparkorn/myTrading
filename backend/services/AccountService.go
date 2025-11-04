package services

import (
	dto "github.com/VarisNithiparkorn/cryptoGraph/backend/DTO"
	"github.com/VarisNithiparkorn/cryptoGraph/backend/entities"
	"github.com/VarisNithiparkorn/cryptoGraph/backend/repositories"
	"github.com/VarisNithiparkorn/cryptoGraph/backend/utils"
)

func CreateAccount(accountDTO dto.Account) {

	account := entities.Account{
		Username: accountDTO.Username,
		Password: utils.Hash(accountDTO.Password),
		IsActive:false,
	}
	repositories.InsertOneAccount(account)
}