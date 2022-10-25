package docmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Document struct {
		ID        primitive.ObjectID `bson:"_id"`
		Drawer    primitive.ObjectID `bson:"drawer"`
		Title     string             `bson:"title"`
		Content   string             `bson:"content"`
		Status    string             `bson:"status"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
		//CreatedBy time.Time          `bson:"createdBy"`
		//UpdatedBy time.Time          `bson:"updatedBy"`
	}
)
