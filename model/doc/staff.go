package docmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Staff struct {
		ID         primitive.ObjectID `bson:"_id"`
		Department primitive.ObjectID `bson:"department"`
		Name       string             `bson:"name"`
		IsRoot     bool               `bson:"isRoot"`
		Phone      string             `bson:"phone""`
		Password   string             `bson:"password"`
		Active     string             `bson:"active"`
		CreatedAt  time.Time          `bson:"createdAt"`
		UpdatedAt  time.Time          `bson:"updatedAt"`
	}
)
