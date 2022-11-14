package docmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	CompanyType struct {
		ID        primitive.ObjectID `bson:"_id"`
		Code      string             `bson:"code"`
		Name      string             `bson:"name"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}
)
