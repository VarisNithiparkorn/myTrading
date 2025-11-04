package controllers

import (
	"fmt"

	dto "github.com/VarisNithiparkorn/cryptoGraph/DTO"
	"github.com/VarisNithiparkorn/cryptoGraph/services"
	"github.com/gin-gonic/gin"
)

	 
func HandleRegister(c *gin.Context) {
	var newAccount dto.Account
	err := c.ShouldBindJSON(&newAccount)
	if(err != nil){
		fmt.Println("BINDING FAILED! Error:", err.Error())
	}
	services.CreateAccount(newAccount)
}