package docmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Company struct {
		ID        primitive.ObjectID `bson:"_id"`
		Type      primitive.ObjectID `bson:"type"`
		Name      string             `bson:"name"`
		Active    string             `bson:"active"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}
)
