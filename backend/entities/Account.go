package entities

import "time"

type Account struct {
	Username  string
	Password  string
	IsActive bool
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time
}