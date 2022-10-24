package docmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Drawer struct {
		ID        primitive.ObjectID `bson:"_id"`
		Cabinet   primitive.ObjectID `bson:"cabinet"`
		Name      string             `bson:"name"`
		Active    string             `bson:"active"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}
)
