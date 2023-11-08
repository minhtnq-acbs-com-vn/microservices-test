package service

import (
	"context"
	"github.com/gookit/goutil/errorx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"microservice-test/book/db"
	"microservice-test/proto/book"
)

type Service struct {
	book.UnimplementedBookServiceServer
	ConnectionString string
}

func New(connectionString string) *Service {
	return &Service{ConnectionString: connectionString}
}

func (s *Service) SaveBooking(ctx context.Context, req *book.BookRequest) (*book.BookResponse, error) {
	client, err := db.New(s.ConnectionString)
	if err != nil {
		return nil, errorx.Wrap(err, "[RPC] SaveBooking New DB Failed")
	}
	result, err := client.Database("book").Collection("book").InsertOne(ctx, req)
	if err != nil {
		return nil, errorx.Wrap(err, "[RPC] SaveBooking InsertOne Failed")
	}
	req.Id = result.InsertedID.(primitive.ObjectID).Hex()

	res := &book.BookResponse{
		Request:    req,
		HelperName: result.InsertedID.(primitive.ObjectID).Hex(),
	}
	return res, nil
}
