package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	docmodel "golangApi/model/doc"
	"golangApi/module/database"
	"log"
)

type CabinetInterface interface {
	Find(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []docmodel.Cabinet)

	FindOne(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc docmodel.Cabinet, err error)

	Create(ctx context.Context, doc interface{}) (err error)

	UpdateOne(ctx context.Context, filter interface{}, doc interface{}, opts ...*options.UpdateOptions) (err error)

	Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (count int64, err error)
}

type cabinetImp struct {
}

func Cabinet() CabinetInterface {
	return cabinetImp{}
}

func (c cabinetImp) Find(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []docmodel.Cabinet) {
	var (
		col = database.CabinetCol()
	)

	cursor, err := col.Find(ctx, cond, opts...)

	if err != nil {
		log.Fatal(err)
		return
	}

	// Close cursor
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &docs); err != nil {
		log.Fatal(err)
	}

	return
}

func (c cabinetImp) FindOne(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc docmodel.Cabinet, err error) {
	var (
		col = database.CabinetCol()
	)

	err = col.FindOne(ctx, cond, opts...).Decode(&doc)

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		//the collection.
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}

	return
}

func (c cabinetImp) Create(ctx context.Context, doc interface{}) (err error) {
	var (
		col = database.CabinetCol()
	)

	_, err = col.InsertOne(ctx, doc)

	if err != nil {
		log.Fatal("err create cabinet")
	}

	return
}

func (c cabinetImp) UpdateOne(ctx context.Context, filter interface{}, doc interface{}, opts ...*options.UpdateOptions) (err error) {
	var (
		col = database.CabinetCol()
	)

	_, err = col.UpdateOne(ctx, filter, bson.D{{"$set", doc}}, opts...)

	if err != nil {
		log.Fatal("err create cabinet")
	}

	return
}

func (c cabinetImp) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (count int64, err error) {
	var (
		col = database.CabinetCol()
	)

	count, err = col.CountDocuments(ctx, filter, opts...)

	if err != nil {
		log.Fatal("err create cabinet")
	}
	return
}
