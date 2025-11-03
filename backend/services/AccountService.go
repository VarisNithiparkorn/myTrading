package services

import (

	dto "github.com/VarisNithiparkorn/cryptoGraph/backend/DTO"
	"github.com/VarisNithiparkorn/cryptoGraph/backend/entities"
	"github.com/VarisNithiparkorn/cryptoGraph/backend/repositories"
)

func CreateAccount(accountDTO dto.Account) {
	account := entities.Account{
		Username: accountDTO.Username,
		Password: accountDTO.Password,
	}
	repositories.InsertOneAccount(account)
}