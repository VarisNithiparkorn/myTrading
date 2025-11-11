package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string `bson:"username" json:"username"`
	Password  string `bson:"password" json:"-"`
	IsActive bool `bson:"is_active" json:"is_active"`
	Role string `bson:"role" json:"role"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}