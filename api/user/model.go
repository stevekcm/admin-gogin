package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auths struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"token"`
	UserAgent string    `json:"user_agent"`
	Device    string    `json:"device"`
}

type User struct {
	ID        primitive.ObjectID `json:"_id"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	Name      string             `json:"name"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Roles     []string           `json:"roles"`
	Auths     []Auths            `json:"auths"`
}
