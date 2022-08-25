package mongo

import (
	"context"
	"time"

	"github.com/nkien0204/projectTemplate/internal/log"
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

type MyCollection struct {
	Collection *mongo.Collection
}

func Init(address string) (*MongoService, error) {
	// address = "mongodb://localhost:27017"
	logger := log.Logger()
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
	logger := log.Logger()
	logger.Warn("close mongo connection", zap.String("address", m.Address))
	defer m.cancel()
	if err := m.Conn.Disconnect(m.ctx); err != nil {
		logger.Error("m.Conn.Disconnect failed", zap.Error(err))
		return err
	}
	return nil
}

func (m *MongoService) GetCollection(databaseName string, collectionName string) *MyCollection {
	return &MyCollection{m.Conn.Database(databaseName).Collection(collectionName)}
}

func (m *MyCollection) InsertOne(document interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.Collection.InsertOne(ctx, document)
	return err
}

func (m *MyCollection) FindOneByObjectId(id primitive.ObjectID) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return m.Collection.FindOne(ctx, bson.M{"_id": id})
}

// MUST handle error
func (m *MyCollection) FindOneByIdString(id string) (*mongo.SingleResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := m.Collection.FindOne(ctx, bson.M{"_id": objectId})
	return result, result.Err()
}

// MUST handle error.
//
// Implement when received cursor:
//
/*	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cursor.Close(ctx)
		cancel()
	}()
	for cursor.Next(ctx) {
		var document models.Test
		err := cursor.Decode(&document)
		if err != nil {
			t.Errorf("%v", err.Error())
			return
		} else {
			result = append(result, document)
		}
	}
	if err := cursor.Err(); err != nil {
		t.Errorf("%v", err)
	} else {
		t.Log("result: ", result)
	}
*/
func (m *MyCollection) Find(filter bson.D, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return m.Collection.Find(ctx, filter, opts...)
}