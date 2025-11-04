package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	bytes,err :=bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if(err != nil){
		log.Println("hash password fail" + err.Error())
	}
	return string(bytes)
}

func Compare(plainText string, hashedPassword string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(plainText))
	return err == nil
}