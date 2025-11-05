package services

import (
	"github.com/VarisNithiparkorn/cryptoGraph/entities"
	"github.com/VarisNithiparkorn/cryptoGraph/repositories"
)

func getAccountByUsername(username string) *entities.Account{
	account := repositories.FindAccountByUsername(username)
	return account
}