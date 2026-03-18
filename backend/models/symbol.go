package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Symbol struct{
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
	Image string `bson:"image" json:"image"`
	CategoryID string `bson:"categoryid" json:"categoryid"`
}