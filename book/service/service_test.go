package service_test

import (
	"context"
	"fmt"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"microservice-test/book/service"
	helperService "microservice-test/helper/service"
	"microservice-test/proto/book"
	"microservice-test/proto/helper"
	"net"
	"os"
	"testing"
)

var dbClient *mongo.Client
var connectionString string
var bookService *service.Service

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pull mongodb docker image for version 5.0
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	connectionString = fmt.Sprintf("mongodb://localhost:%s", resource.GetPort("27017/tcp"))
	StartHelperServiceForTesting()

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	err = pool.Retry(func() error {
		var err error
		dbClient, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				connectionString,
			),
		)
		if err != nil {
			return err
		}
		return dbClient.Ping(context.TODO(), nil)
	})

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	bookService = service.New(connectionString)

	// run tests
	code := m.Run()

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	// disconnect mongodb client
	if err = dbClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

	os.Exit(code)
}

func TestSaveBooking(t *testing.T) {
	req := &book.BookRequest{
		From: "customer",
		Desc: "clean",
	}
	res, err := bookService.SaveBooking(context.Background(), req)

	assert.Nil(t, err)
	assert.Equal(t, res.Request.From, req.From)
	assert.Equal(t, res.Request.Desc, req.Desc)

	objectId, err := primitive.ObjectIDFromHex(res.Request.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, objectId)

	// check if data is saved in mongodb
	collection := dbClient.Database("book").Collection("book")
	var result book.BookRequest
	err = collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, result.From, req.From)
	assert.Equal(t, result.Desc, req.Desc)
}

func StartHelperServiceForTesting() {
	lis, err := net.Listen("tcp", ":12000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	helperS := helperService.New(connectionString)

	helper.RegisterHelperServer(grpcServer, helperS)

	go func() {
		log.Println("Server Helper is listening on port: 12000")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
