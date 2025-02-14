package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate: "required",min=2,max=50"`
	Last_name     *string            `json:"last_name" validate: "required",min=2,max=50"`
	Password      *string            `json: "password" validate: "required,min=6"`
	Email         *string            `json: "email" validate: "email, required"`
	Token         *string            `json: "token"`
	Refresh_token *string            `json: "refrsh_token"`
	User_type     *string            `json: "user_type" validate: "required,eq=ADMIN|eq=USER"`
	Created_at    time.Time          `json: "created_at"`
	Updated_at    time.Time          `json: "updated_at"`
	User_id       string             `json: "user_id"`
}

type UserSettings struct {
	ID               primitive.ObjectID `bson:"_id"`
	Languages        []string           `bson:"languages,omitempty"`
	Default_language string             `bson:"default_language,omitempty"`
	User_id          string             `bson:user_id,omitempty`
}
