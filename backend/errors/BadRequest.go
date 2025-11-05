package errorsHandle

import (

	"time"
)
 

type BadRequestErr struct {
	Status    int       `json:"status"`
	Message   string    `json:"message"`
	TimeStamp time.Time `json:"timestamp"` 
}

func (err *BadRequestErr) Error() string{
	return err.Message
}

func CreateBadRequest(status int,message string) *BadRequestErr{
	return &BadRequestErr{
		Status: status,
		Message: message,
		TimeStamp: time.Now(),
	}
}