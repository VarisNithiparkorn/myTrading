package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	dto "github.com/VarisNithiparkorn/cryptoGraph/DTO"
	errorsHandle "github.com/VarisNithiparkorn/cryptoGraph/errors"
	"github.com/VarisNithiparkorn/cryptoGraph/services"
	"github.com/gin-gonic/gin"
)

func ErrorResponse(e error,c *gin.Context){
	var badRequestErr *errorsHandle.BadRequestErr
    var conflictErr *errorsHandle.ConflictErr
    var internalServerErr *errorsHandle.InternalServerErr
    var unauthorizedErr *errorsHandle.UnAuthorizedErr


    if errors.As(e, &badRequestErr) {

        c.JSON(http.StatusBadRequest, badRequestErr)
        
    } else if errors.As(e, &conflictErr) {
        c.JSON(http.StatusConflict, conflictErr)
        
    } else if errors.As(e, &internalServerErr) {


        c.JSON(http.StatusInternalServerError, internalServerErr)
        
    } else if errors.As(e, &unauthorizedErr) {
        c.JSON(http.StatusUnauthorized, unauthorizedErr)
        
    } else {
        log.Printf("UNHANDLED ERROR: %v", e)
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  http.StatusInternalServerError,
            "message": "An unexpected server error occurred.",
        })
    }
}


func HandleRegister(c *gin.Context) {
	var newAccount dto.Account
	err := c.ShouldBindJSON(&newAccount)
	if(err != nil){
		fmt.Println("BINDING FAILED! Error:", err.Error())
	}
	httpErr := services.CreateAccount(newAccount)
	if(httpErr != nil){
		ErrorResponse(httpErr,c)
		return
	}
	c.JSON(http.StatusCreated,gin.H{"message":"created success"})
}
func HandleVerifyEmail(c *gin.Context) {
	token := c.Query("token") 

	err := services.VerifyEmail(token)
	if(err != nil){
		ErrorResponse(err,c)
	}
	c.JSON(http.StatusOK,"")
}