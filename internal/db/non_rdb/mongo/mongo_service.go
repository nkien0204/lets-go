package mongo

import (
	"context"
	"time"

	"github.com/nkien0204/projectTemplate/internal/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoService struct {
	Address string
	Conn *mongo.Client
}

func Init(address string) MongoService {
	// address = "mongodb://localhost:27017"
	logger := log.Logger()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		logger.Fatal("mongo.Connect failed", zap.Error(err))
	}
	return MongoService{
		Address: address,
		Conn: client,
	}
}