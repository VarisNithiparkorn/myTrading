package errorsHandle

import "time"

type ConflictErr struct {
	Status    int       `json:"status"`    
	Message   string    `json:"message"`   
	TimeStamp time.Time `json:"timestamp"` 
}

func (err *ConflictErr) Error() string {
	return err.Message
}

func CreateConflictErr(status int, message string) *ConflictErr {
	return &ConflictErr{
		Status:    status,
		Message:   message,
		TimeStamp: time.Now(),
	}
}

