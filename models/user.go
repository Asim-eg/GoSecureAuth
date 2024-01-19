package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"id" json:"id,omitempty"`
	FirstName    *string            `bson:"firstName" json:"firstName" validate:"required" min:"2" max:"100"`
	LastName     *string            `bson:"lastName" json:"lastName" validate:"required" min:"2" max:"100"`
	Email        *string            `bson:"email" json:"email" validate:"required,email"`
	Password     *string            `bson:"password" json:"password" validate:"required,min=6"`
	Phone        *string            `bson:"phone" json:"phone" validate:"required"`
	Token        *string            `bson:"token" json:"token"`
	UserType     *string            `bson:"userType" json:"userType" validate:"required, eq=ADMIN|eq=USER"`
	RefreshToken *string            `bson:"refreshToken" json:"refreshToken"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
	UserID       string             `bson:"userId" json:"userId"`
}
