package entities

import "time"

type Account struct {
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}