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

type DocumentInterface interface {
	Find(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []docmodel.Document)

	FindOne(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc docmodel.Document, err error)

	Create(ctx context.Context, doc interface{}) (err error)

	UpdateOne(ctx context.Context, filter interface{}, doc interface{}, opts ...*options.UpdateOptions) (err error)

	Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (count int64, err error)
}

type documentImp struct {
}

func Document() DocumentInterface {
	return documentImp{}
}

func (c documentImp) Find(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []docmodel.Document) {
	var (
		col = database.DocumentCol()
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

func (c documentImp) FindOne(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc docmodel.Document, err error) {
	var (
		col = database.DocumentCol()
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

func (c documentImp) Create(ctx context.Context, doc interface{}) (err error) {
	var (
		col = database.DocumentCol()
	)

	_, err = col.InsertOne(ctx, doc)

	if err != nil {
		log.Fatal("err create document")
	}

	return
}

func (c documentImp) UpdateOne(ctx context.Context, filter interface{}, doc interface{}, opts ...*options.UpdateOptions) (err error) {
	var (
		col = database.DocumentCol()
	)

	_, err = col.UpdateOne(ctx, filter, bson.D{{"$set", doc}}, opts...)

	if err != nil {
		log.Fatal("err create document")
	}

	return
}

func (c documentImp) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (count int64, err error) {
	var (
		col = database.DocumentCol()
	)

	count, err = col.CountDocuments(ctx, filter, opts...)

	if err != nil {
		log.Fatal("err create document")
	}
	return
}
