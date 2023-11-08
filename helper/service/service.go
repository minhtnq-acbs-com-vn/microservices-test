package service

import (
	"context"
	"fmt"
	"github.com/gookit/goutil/errorx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"microservice-test/book/db"
	"microservice-test/proto/book"
	"microservice-test/proto/helper"
)

type Service struct {
	helper.UnimplementedHelperServer
	ConnectionString string
}

func New(connectionString string) *Service {
	return &Service{ConnectionString: connectionString}
}

func (s *Service) UpdateJob(ctx context.Context, req *book.BookRequest) (*book.BookResponse, error) {
	client, err := db.New(s.ConnectionString)
	if err != nil {
		return nil, errorx.Wrap(err, fmt.Sprintf("[DB] Connect Failed %v", err))
	}

	helperCollection := client.Database("book").Collection("helper")
	if err := SetupFakeData(helperCollection); err != nil {
		return nil, errorx.Wrap(err, "[RPC] UpdateJob SetupFakeData Failed")
	}

	numDocs, err := helperCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return nil, errorx.Wrap(err, "[RPC] UpdateJob CountDocuments Failed")
	}

	cursor, err := helperCollection.Find(context.Background(), bson.M{}, options.Find().SetLimit(1).SetSkip(int64(rand.Intn(int(numDocs)))))
	if err != nil {
		return nil, errorx.Wrap(err, "[RPC] UpdateJob Find Failed")
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, errorx.Wrap(err, "[RPC] UpdateJob Cursor.All Failed")
	}

	randomHelper := results[0]["name"].(string)

	objectId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, errorx.Wrap(err, "[RPC] UpdateJob ObjectIDFromHex Failed")
	}

	if _, err := helperCollection.UpdateByID(ctx, objectId, bson.M{"$set": bson.M{"name": randomHelper}}); err != nil {
		return nil, errorx.Wrap(err, "[RPC] UpdateJob UpdateByID Failed")
	}

	res := &book.BookResponse{
		Request:    req,
		HelperName: randomHelper,
	}
	return res, nil
}

func SetupFakeData(helperCollection *mongo.Collection) error {
	helperData := []interface{}{
		bson.D{{"name", "Helper 1"}, {"phone", 0}},
		bson.D{{"name", "Helper 2"}, {"phone", 1}},
		bson.D{{"name", "Helper 3"}, {"phone", 2}},
	}
	results, err := helperCollection.InsertMany(context.Background(), helperData)
	if err != nil {
		return errorx.Wrap(err, "[RPC] SetupFakeData InsertFakeData Failed")
	}
	fmt.Println(results.InsertedIDs)
	return nil
}
