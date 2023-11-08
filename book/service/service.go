package service

import (
	"context"
	"fmt"
	"github.com/gookit/goutil/errorx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	helperClient "microservice-test/book/client"
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
	fmt.Println("[RPC] SaveBooking Called with req: ", req)
	client, err := db.New(s.ConnectionString)
	if err != nil {
		return nil, errorx.Wrap(err, "[RPC] SaveBooking New DB Failed")
	}
	result, err := client.Database("book").Collection("book").InsertOne(ctx, req)
	if err != nil {
		return nil, errorx.Wrap(err, "[RPC] SaveBooking InsertOne Failed")
	}
	req.Id = result.InsertedID.(primitive.ObjectID).Hex()

	res, err := helperClient.CallHelperToUpdate(req)
	if err != nil {
		return nil, errorx.Wrap(err, "[RPC] SaveBooking CallHelperToUpdate Failed")
	}
	fmt.Println("[RPC] SaveBooking Called with res: ", res)

	return res, nil
}
