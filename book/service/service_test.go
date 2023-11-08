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
	"log"
	"microservice-test/book/service"
	"microservice-test/proto/book"
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
		Env: []string{
			// username and password for mongodb superuser
			"MONGO_INITDB_ROOT_USERNAME=root",
			"MONGO_INITDB_ROOT_PASSWORD=password",
		},
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

	connectionString = fmt.Sprintf("mongodb://root:password@localhost:%s", resource.GetPort("27017/tcp"))

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
		From:   "customer",
		Helper: "",
		Desc:   "clean",
	}
	res, err := bookService.SaveBooking(context.Background(), req)

	assert.Nil(t, err)
	assert.Equal(t, res.Status, "success")
	assert.NotEmpty(t, res.Id)

	objectId, err := primitive.ObjectIDFromHex(res.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, objectId)

	// check if data is saved in mongodb
	collection := dbClient.Database("book").Collection("book")
	var result book.BookRequest
	err = collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, result.From, req.From)
	assert.Equal(t, result.Helper, req.Helper)
	assert.Equal(t, result.Desc, req.Desc)
}
