package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	dto "github.com/VarisNithiparkorn/cryptoGraph/DTO"
	errorsHandle "github.com/VarisNithiparkorn/cryptoGraph/errors"
	"github.com/VarisNithiparkorn/cryptoGraph/services"
	"github.com/VarisNithiparkorn/cryptoGraph/utils"
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
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  http.StatusInternalServerError,
            "message": "An unexpected server error occurred.",
        })
    }
}
var
( 
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	specialChars = "!@#$%^&*-_+="
)
                     

func validateFormatPassword(password string) bool{
	haveUpper := false
	haveLower := false
	haveDigit := false
	for _,char := range password{
		if haveUpper && haveLower && haveDigit {
			break
		}
		if !haveUpper {
			haveUpper = unicode.IsUpper(char)	
		}
		if  !haveLower{
			haveLower = unicode.IsLower(char)
		}
		if  !haveDigit{
			haveDigit = unicode.IsDigit(char)
		}
	}
	haveSpecialCharacter := strings.ContainsAny(password,specialChars)
	if haveUpper && haveLower && haveDigit && haveSpecialCharacter{
		return true
	}
	return false
}
func validateEmail(username string)error{
	if utils.IsemptyField(username){
		return errorsHandle.CreateBadRequest(400,"username is empty")
	}else if !emailRegex.MatchString(username){
		return errorsHandle.CreateBadRequest(400,"username is invalid format")
	} 
	return nil
}
func validateAccount(account *dto.Account) error{
	username := account.Username
	password := account.Password
	role := account.Role
	err := validateEmail(username)
	if err != nil{
		return err
	}
	if utils.IsemptyField(password){
		return errorsHandle.CreateBadRequest(400,"password is empty")
	} else if  len(password) < 12 {
		return errorsHandle.CreateBadRequest(400,"required password length more than 12")
	}else if !validateFormatPassword(password) {
		return errorsHandle.CreateBadRequest(400,"password is invalid format")
	}
	if utils.IsemptyField(role) {
		return errorsHandle.CreateBadRequest(400,"role is empty")
	}else if role != "user" && role != "admin" {
		return errorsHandle.CreateBadRequest(400,"invalid role")
	}
	return nil
}

func HandleRegister(c *gin.Context) {
	var newAccount dto.Account
	err := c.ShouldBindJSON(&newAccount)
	if(err != nil){
		fmt.Println("BINDING FAILED! Error:", err.Error())
	}
	err2 := validateAccount(&newAccount)
	if err2 != nil{
		ErrorResponse(err2,c)
		return
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
		return
	}
	c.JSON(http.StatusOK,"")
}

func HandleLogin(c *gin.Context){
	var newAccount dto.Account
	err := c.ShouldBindJSON(&newAccount)
	if(err != nil){
		fmt.Println("BINDING FAILED! Error:", err.Error())
	}
	valid := emailRegex.MatchString(newAccount.Username)
	if !valid{
		ErrorResponse(err,c)
		return
	}
	token ,err2 := services.Login(&newAccount)
	if err2 != nil{
		ErrorResponse(err2,c)
		return
	}
	c.JSON(200,gin.H{"accesstoken":token})
}