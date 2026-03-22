package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Phrase struct{
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID primitive.ObjectID `bson:"user_id" json:"user_id"`
	Symbols []primitive.ObjectID `bson:"symbols" json:"symbols"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type PhraseRequest struct {
	UserID  string   `json:"user_id"`
	Symbols []string `json:"symbols"`
}