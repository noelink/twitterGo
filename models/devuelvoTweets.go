package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DevuelvoTweets struct {
	ID      primitive.ObjectID `bson:"id" json:"id,omitempty"`
	UserId  string             `bson:"userid " json:"userid"`
	Mensaje string             `bson:"mensaje " json:"mensaje"`
	Fecha   time.Time          `bson:"fecha" json:"fecha"`
}
