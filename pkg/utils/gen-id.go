package utils

import (
	uuid "github.com/satori/go.uuid"
	bsonp "go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUUID() string { //TODO: faster
	return uuid.NewV4().String()
}

func NewOID() string {
	return bsonp.NewObjectID().String()
}
