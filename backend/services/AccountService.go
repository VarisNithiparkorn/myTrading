package services

import (
	dto "github.com/VarisNithiparkorn/cryptoGraph/DTO"
	"github.com/VarisNithiparkorn/cryptoGraph/entities"
	"github.com/VarisNithiparkorn/cryptoGraph/repositories"
	"github.com/VarisNithiparkorn/cryptoGraph/utils"
)

func CreateAccount(accountDTO dto.Account) {

	account := entities.Account{
		Username: accountDTO.Username,
		Password: utils.Hash(accountDTO.Password),
		IsActive:false,
	}
	repositories.InsertOneAccount(account)
}