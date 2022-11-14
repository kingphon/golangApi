package docmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Permission struct {
		ID   primitive.ObjectID `bson:"_id"`
		Name string             `bson:"name"`
		Code string             `bson:"code"`
	}
)
