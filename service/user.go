package service

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	mongo.Connect(context.Background(), mongo.ClientOptions{}.ApplyURI("mongodb://localhost:27017"))
}
