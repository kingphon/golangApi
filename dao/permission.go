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

type PermissionInterface interface {
	Find(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []docmodel.Permission)

	FindOne(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc docmodel.Permission, err error)

	Create(ctx context.Context, doc interface{}) (err error)

	UpdateOne(ctx context.Context, filter interface{}, doc interface{}, opts ...*options.UpdateOptions) (err error)

	Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (count int64, err error)
}

type permissionImp struct {
}

func Permission() PermissionInterface {
	return permissionImp{}
}

func (p permissionImp) Find(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []docmodel.Permission) {
	var (
		col = database.PermissionCol()
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

func (p permissionImp) FindOne(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc docmodel.Permission, err error) {
	var (
		col = database.PermissionCol()
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

func (p permissionImp) Create(ctx context.Context, doc interface{}) (err error) {
	var (
		col = database.PermissionCol()
	)

	_, err = col.InsertOne(ctx, doc)

	if err != nil {
		log.Fatal("err create permission")
	}

	return
}

func (p permissionImp) UpdateOne(ctx context.Context, filter interface{}, doc interface{}, opts ...*options.UpdateOptions) (err error) {
	var (
		col = database.PermissionCol()
	)

	_, err = col.UpdateOne(ctx, filter, bson.D{{"$set", doc}}, opts...)

	if err != nil {
		log.Fatal("err create permission")
	}

	return
}

func (p permissionImp) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (count int64, err error) {
	var (
		col = database.PermissionCol()
	)

	count, err = col.CountDocuments(ctx, filter, opts...)

	if err != nil {
		log.Fatal("err create permission")
	}
	return
}
