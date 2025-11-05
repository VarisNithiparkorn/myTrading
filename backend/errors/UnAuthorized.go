package errorsHandle

import (

	"time"
)
 

type UnAuthorizedErr struct {
	Status    int       `json:"status"`    
	Message   string    `json:"message"`   
	TimeStamp time.Time `json:"timestamp"` 
}

func (err *UnAuthorizedErr) Error() string{
	return err.Message
}

func CreateUnAuthorizedErr(status int,message string) *UnAuthorizedErr{
	return &UnAuthorizedErr{
		Status: status,
		Message: message,
		TimeStamp: time.Now(),
	}
}