package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection *mongo.Collection
	validate       = validator.New()
)

func MatchUserTypeToUid(c *gin.Context, uid string) error {
	// Get the user type from the JWT
	userType := c.GetString("userType")
	id := c.GetString("uid")

	if userType == "USER" && id != uid {
		return errors.New("Unauthorized Access")
	}

	err := CheckUserType(c, userType)
	return err
}

func CheckUserType(c *gin.Context, role string) error {
	userType := c.GetString("userType")

	if userType != role {
		return errors.New("Unauthorized Access")
	}
	return nil
}
