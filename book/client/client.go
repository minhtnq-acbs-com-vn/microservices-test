package client

import (
	"context"
	"fmt"
	"github.com/gookit/goutil/errorx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"microservice-test/proto/book"
	"microservice-test/proto/helper"
)

const helperConnectionString = "host.docker.internal:12000"

func CallHelperToUpdate(req *book.BookRequest) (*book.BookResponse, error) {
	fmt.Println("[CLIENT] calling helper to update job with request: ", req)
	conn, err := grpc.Dial(helperConnectionString, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errorx.Wrap(err, "[CLIENT] did not connect")
	}
	defer conn.Close()

	helperClient := helper.NewHelperClient(conn)

	res, err := helperClient.UpdateJob(context.Background(), req)
	if err != nil {
		return nil, errorx.Wrap(err, "[CLIENT] could not update job")
	}

	fmt.Println("[CLIENT] got response from helper: ", res)
	return res, nil
}
