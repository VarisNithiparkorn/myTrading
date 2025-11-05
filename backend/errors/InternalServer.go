package errorsHandle

import (

	"time"
)
type InternalServerErr struct {
	Status    int       `json:"status"`    
	Message   string    `json:"message"`   
	TimeStamp time.Time `json:"timestamp"` 
}

func (err *InternalServerErr) Error() string{
	return err.Message
}

func CreateInternalServerErr(status int,message string) *ConflictErr{
	return &ConflictErr{
		Status: status,
		Message: message,
		TimeStamp: time.Now(),
	}
}



