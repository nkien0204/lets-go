package test

import (
	"testing"

	"github.com/nkien0204/lets-go/internal/infrastructure/db/non_rdb/mongo"
	"github.com/nkien0204/lets-go/internal/infrastructure/db/non_rdb/mongo/models"
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

	collectionInterface, err := mongoService.GetCollection(DatabaseName, models.TestCollectionName)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	collection, ok := collectionInterface.(*mongo.MyCollection[models.Test])
	if !ok {
		t.Errorf("failed")
		return
	}

	result, err := collection.Find(bson.D{})
	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}
	t.Log(result)
}

func TestFindOneById(t *testing.T) {
	mongoService, err := mongo.Init("mongodb://admin:admin@localhost:27017")
	if err != nil {
		t.Errorf("mongo.Init failed %v", err.Error())
		return
	}

	collectionInterface, err := mongoService.GetCollection(DatabaseName, models.TestCollectionName)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	collection, ok := collectionInterface.(*mongo.MyCollection[models.Test])
	if !ok {
		t.Errorf("failed")
		return
	}
	objectID, err := primitive.ObjectIDFromHex("6306e0c436c5618ce062355d")
	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}
	result, err := collection.FindOneByObjectId(objectID)
	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}
	t.Log(result)
}

func TestFindOneByIdString(t *testing.T) {
	mongoService, err := mongo.Init("mongodb://admin:admin@localhost:27017")
	if err != nil {
		t.Errorf("mongo.Init failed %v", err.Error())
		return
	}
	collectionInterface, err := mongoService.GetCollection(DatabaseName, models.TestCollectionName)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	collection, ok := collectionInterface.(*mongo.MyCollection[models.Test])
	if !ok {
		t.Errorf("failed")
		return
	}
	result, err := collection.FindOneByIdString("6306e0c436c5618ce062355d")
	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}
	t.Log(result)
}

func TestInsertMany(t *testing.T) {
	mongoService, err := mongo.Init("mongodb://admin:admin@localhost:27017")
	if err != nil {
		t.Errorf("mongo.Init failed %v", err.Error())
		return
	}
	collectionInterface, err := mongoService.GetCollection(DatabaseName, models.TestCollectionName)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	collection, ok := collectionInterface.(*mongo.MyCollection[models.Test])
	if !ok {
		t.Errorf("failed")
		return
	}
	update := []models.Test{
		{
			Id:           primitive.NewObjectID(),
			AnotherField: "test3",
			MyField:      "TEST3",
		},
		{
			Id:           primitive.NewObjectID(),
			AnotherField: "test4",
			MyField:      "TEST4",
		},
	}
	if err := collection.InsertMany(update); err != nil {
		t.Errorf("%v", err.Error())
		return
	}
}

func TestUpdateOne(t *testing.T) {
	mongoService, err := mongo.Init("mongodb://admin:admin@localhost:27017")
	if err != nil {
		t.Errorf("mongo.Init failed %v", err.Error())
		return
	}
	collectionInterface, err := mongoService.GetCollection(DatabaseName, models.TestCollectionName)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	collection, ok := collectionInterface.(*mongo.MyCollection[models.Test])
	if !ok {
		t.Errorf("failed")
		return
	}
	err = collection.UpdateOne(bson.M{"anotherfield": "test44"}, bson.M{"anotherfield": "test4411", "myfield": "kien nguyen1231233"})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
}

func TestUpdateMany(t *testing.T) {
	mongoService, err := mongo.Init("mongodb://admin:admin@localhost:27017")
	if err != nil {
		t.Errorf("mongo.Init failed %v", err.Error())
		return
	}
	collectionInterface, err := mongoService.GetCollection(DatabaseName, models.TestCollectionName)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	collection, ok := collectionInterface.(*mongo.MyCollection[models.Test])
	if !ok {
		t.Errorf("failed")
		return
	}
	err = collection.UpdateMany(bson.M{"anotherfield": "test4411"}, bson.M{"anotherfield": "123test441", "myfield": "123kien nguyen123"})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
}

func TestDeleteOne(t *testing.T) {
	mongoService, err := mongo.Init("mongodb://admin:admin@localhost:27017")
	if err != nil {
		t.Errorf("mongo.Init failed %v", err.Error())
		return
	}
	collectionInterface, err := mongoService.GetCollection(DatabaseName, models.TestCollectionName)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	collection, ok := collectionInterface.(*mongo.MyCollection[models.Test])
	if !ok {
		t.Errorf("failed")
		return
	}
	err = collection.DeleteOne(bson.M{"anotherfield": "TEST1"})
	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}
}

func TestDeleteMany(t *testing.T) {
	mongoService, err := mongo.Init("mongodb://admin:admin@localhost:27017")
	if err != nil {
		t.Errorf("mongo.Init failed %v", err.Error())
		return
	}
	collectionInterface, err := mongoService.GetCollection(DatabaseName, models.TestCollectionName)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	collection, ok := collectionInterface.(*mongo.MyCollection[models.Test])
	if !ok {
		t.Errorf("failed")
		return
	}
	err = collection.DeleteMany(bson.M{"anotherfield": "123test441"})
	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}
}
