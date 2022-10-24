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

type DrawerInterface interface {
	Find(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []docmodel.Drawer)

	FindOne(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc docmodel.Drawer, err error)

	Create(ctx context.Context, doc interface{}) (err error)

	UpdateOne(ctx context.Context, filter interface{}, doc interface{}, opts ...*options.UpdateOptions) (err error)

	Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (count int64, err error)
}

type drawerImp struct {
}

func Drawer() DrawerInterface {
	return drawerImp{}
}

func (c drawerImp) Find(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []docmodel.Drawer) {
	var (
		col = database.DrawerCol()
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

func (c drawerImp) FindOne(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc docmodel.Drawer, err error) {
	var (
		col = database.DrawerCol()
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

func (c drawerImp) Create(ctx context.Context, doc interface{}) (err error) {
	var (
		col = database.DrawerCol()
	)

	_, err = col.InsertOne(ctx, doc)

	if err != nil {
		log.Fatal("err create drawer")
	}

	return
}

func (c drawerImp) UpdateOne(ctx context.Context, filter interface{}, doc interface{}, opts ...*options.UpdateOptions) (err error) {
	var (
		col = database.DrawerCol()
	)

	_, err = col.UpdateOne(ctx, filter, bson.D{{"$set", doc}}, opts...)

	if err != nil {
		log.Fatal("err create drawer")
	}

	return
}

func (c drawerImp) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (count int64, err error) {
	var (
		col = database.DrawerCol()
	)

	count, err = col.CountDocuments(ctx, filter, opts...)

	if err != nil {
		log.Fatal("err create drawer")
	}
	return
}
