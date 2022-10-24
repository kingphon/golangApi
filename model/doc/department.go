package docmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Department struct {
		ID         primitive.ObjectID `bson:"_id"`
		Company    primitive.ObjectID `bson:"company"`
		Name       string             `bson:"name"`
		Active     string             `bson:"active"`
		Permission []string           `bson:"permission"`
		CreatedAt  time.Time          `bson:"createdAt"`
		UpdatedAt  time.Time          `bson:"updatedAt"`
	}
)
