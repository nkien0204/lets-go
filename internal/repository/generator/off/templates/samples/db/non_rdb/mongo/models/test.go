package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const TestCollectionName = "test"

type Test struct {
	Id           primitive.ObjectID `bson:"_id"`
	AnotherField string             `bson:"anotherfield"`
	MyField      string             `bson:"myfield"`
}
