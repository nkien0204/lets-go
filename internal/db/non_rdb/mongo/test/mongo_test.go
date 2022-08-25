package test

import (
	"context"
	"testing"
	"time"

	"github.com/nkien0204/projectTemplate/internal/db/non_rdb/mongo"
	"github.com/nkien0204/projectTemplate/internal/db/non_rdb/mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const DatabaseName = "test"

func TestInitConnection(t *testing.T) {
	mongoService, err := mongo.Init("mongodb://localhost:27017")
	if err != nil {
		t.Errorf("mongo.Init failed %v", err.Error())
	} else {
		t.Log("mongo.Init successfully", mongoService.Address)
	}
}

func TestCloseConnection(t *testing.T) {
	mongoService, err := mongo.Init("mongodb://localhost:27017")
	if err != nil {
		t.Errorf("mongo.Init failed %v", err.Error())
	} else {
		t.Log("mongo.Init successfully", mongoService.Address)
		if err := mongoService.Close(); err != nil {
			t.Errorf("mongoService.Close failed %v", err.Error())
		} else {
			t.Log("mongoService.Close OK")
		}
	}
}

func TestFind(t *testing.T) {
	mongoService, err := mongo.Init("mongodb://admin:admin@localhost:27017")
	if err != nil {
		t.Errorf("mongo.Init failed %v", err.Error())
		return
	}
	result := make([]models.Test, 0)
	collection := mongoService.GetCollection(DatabaseName, models.TestCollectionName)
	cursor, err := collection.Find(bson.D{})
	if err != nil {
		t.Errorf("collection.Find failed %v", err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
}

func TestFindOneById(t *testing.T) {
	mongoService, err := mongo.Init("mongodb://admin:admin@localhost:27017")
	if err != nil {
		t.Errorf("mongo.Init failed %v", err.Error())
	} else {
		collection := mongoService.GetCollection(DatabaseName, models.TestCollectionName)
		objectID, err := primitive.ObjectIDFromHex("6306e0c436c5618ce062355d")
		if err != nil {
			t.Errorf("%v", err.Error())
		} else {
			var result models.Test
			raw := collection.FindOneByObjectId(objectID)
			if err := raw.Decode(&result); err != nil {
				t.Errorf("%v", err.Error())
			} else {
				t.Log(result)
			}
		}
	}
}

func TestFindOneByIdString(t *testing.T) {
	mongoService, err := mongo.Init("mongodb://admin:admin@localhost:27017")
	if err != nil {
		t.Errorf("mongo.Init failed %v", err.Error())
	} else {
		collection := mongoService.GetCollection(DatabaseName, models.TestCollectionName)
		var result models.Test
		raw, err := collection.FindOneByIdString("6306e0c436c5618ce062351d")
		if err != nil {
			t.Errorf("%v", err.Error())
		} else {
			if err := raw.Decode(&result); err != nil {
				t.Errorf("%v", err.Error())
			} else {
				t.Log(result)
			}
		}
	}
}