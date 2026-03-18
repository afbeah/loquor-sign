package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Phrase struct{
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID string `bson:"user_id" json:"user_id"`
	Symbols []string `bson:"symbols" json:"symbols"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}