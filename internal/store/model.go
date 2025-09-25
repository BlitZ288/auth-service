package store

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
	Roles        []string           `bson:"roles"`
	IsVerified   bool               `bson:"is_verified"`
	CreatedAt    time.Time          `bson:"created_at"`
	LastLogin    *time.Time         `bson:"last_login"`
}
