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

type DepartmentInterface interface {
	Find(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []docmodel.Department)

	FindOne(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc docmodel.Department, err error)

	Create(ctx context.Context, doc interface{}) (err error)

	UpdateOne(ctx context.Context, filter interface{}, doc interface{}, opts ...*options.UpdateOptions) (err error)

	Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (count int64, err error)
}

type departmentImp struct {
}

func Department() DepartmentInterface {
	return departmentImp{}
}

func (c departmentImp) Find(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []docmodel.Department) {
	var (
		col = database.DepartmentCol()
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

func (c departmentImp) FindOne(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc docmodel.Department, err error) {
	var (
		col = database.DepartmentCol()
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

func (c departmentImp) Create(ctx context.Context, doc interface{}) (err error) {
	var (
		col = database.DepartmentCol()
	)

	_, err = col.InsertOne(ctx, doc)

	if err != nil {
		log.Fatal("err create department")
	}

	return
}

func (c departmentImp) UpdateOne(ctx context.Context, filter interface{}, doc interface{}, opts ...*options.UpdateOptions) (err error) {
	var (
		col = database.DepartmentCol()
	)

	_, err = col.UpdateOne(ctx, filter, bson.D{{"$set", doc}}, opts...)

	if err != nil {
		log.Fatal("err create department")
	}

	return
}

func (c departmentImp) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (count int64, err error) {
	var (
		col = database.DepartmentCol()
	)

	count, err = col.CountDocuments(ctx, filter, opts...)

	if err != nil {
		log.Fatal("err create department")
	}
	return
}
