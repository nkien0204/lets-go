package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/nkien0204/lets-go/internal/db/non_rdb/mongo/models"
	"github.com/nkien0204/rolling-logger/rolling"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type MongoService struct {
	ctx     context.Context
	cancel  context.CancelFunc
	Address string
	Conn    *mongo.Client
}

type MyCollection[T any] struct {
	collection *mongo.Collection
}

func Init(address string) (*MongoService, error) {
	// address = "mongodb://user:pass@localhost:27017"
	logger := rolling.New()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		cancel()
		logger.Error("mongo.NewClient failed", zap.Error(err))
		return nil, err
	}
	if err := client.Connect(ctx); err != nil {
		logger.Error("client.Connect failed", zap.Error(err))
		cancel()
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		cancel()
		logger.Error("client.Ping failed", zap.Error(err))
		return nil, err
	}

	return &MongoService{
		Address: address,
		Conn:    client,
		ctx:     ctx,
		cancel:  cancel,
	}, nil
}

func (m *MongoService) Close() error {
	logger := rolling.New()
	logger.Warn("close mongo connection", zap.String("address", m.Address))
	defer m.cancel()
	if err := m.Conn.Disconnect(m.ctx); err != nil {
		logger.Error("m.Conn.Disconnect failed", zap.Error(err))
		return err
	}
	return nil
}

// MUST handle error!
//
//Convert return interface into specific model ("models.Test" for example):
//
/*
	collectionInterface, err := mongoService.GetCollection(DatabaseName, models.TestCollectionName)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	collection := collectionInterface.(*mongo.MyCollection[models.Test])
*/
func (m *MongoService) GetCollection(databaseName string, collectionName string) (interface{}, error) {
	logger := rolling.New()
	switch collectionName {
	case models.TestCollectionName:
		return &MyCollection[models.Test]{
			collection: m.Conn.Database(databaseName).Collection(collectionName),
		}, nil
	default:
		logger.Error("collection name not found", zap.String("name", collectionName))
		return nil, errors.New("collection name not found")
	}
}

// MUST handle error!
func (m *MyCollection[T]) InsertOne(document T) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.collection.InsertOne(ctx, document)
	return err
}

// MUST handle error!
func (m *MyCollection[T]) InsertMany(document []T) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dst := make([]interface{}, len(document))
	for i := range document {
		dst[i] = document[i]
	}
	_, err := m.collection.InsertMany(ctx, dst)
	return err
}

// MUST handle error!
func (m *MyCollection[T]) FindOneByObjectId(id primitive.ObjectID) (result T, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	return result, err
}

// MUST handle error!
func (m *MyCollection[T]) FindOneByIdString(id string) (result T, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	err = m.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)
	return
}

// MUST handle error!
func (m *MyCollection[T]) Find(filter bson.D, opts ...*options.FindOptions) (result []T, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := m.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	result = make([]T, 0)
	for cursor.Next(ctx) {
		var document T
		err = cursor.Decode(&document)
		if err != nil {
			return
		} else {
			result = append(result, document)
		}
	}
	err = cursor.Err()
	return
}

// MUST handle error!
func (m *MyCollection[T]) UpdateOne(filter bson.M, update bson.M, opts ...*options.UpdateOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.collection.UpdateOne(ctx, filter, bson.M{"$set": update}, opts...)
	return err
}

// MUST handle error!
func (m *MyCollection[T]) UpdateMany(filter bson.M, update bson.M, opts ...*options.UpdateOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.collection.UpdateMany(ctx, filter, bson.M{"$set": update}, opts...)
	return err
}

func (m *MyCollection[T]) DeleteOne(filter bson.M, opts ...*options.DeleteOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.collection.DeleteOne(ctx, filter, opts...)
	return err
}

func (m *MyCollection[T]) DeleteMany(filter bson.M, opts ...*options.DeleteOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.collection.DeleteMany(ctx, filter, opts...)
	return err
}
