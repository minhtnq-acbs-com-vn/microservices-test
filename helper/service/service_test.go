package service_test

import (
	"context"
	"fmt"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"microservice-test/helper/service"
	"microservice-test/proto/book"

	"os"
	"testing"
)

var dbClient *mongo.Client
var connectionString string
var helperService *service.Service

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

	helperService = service.New(connectionString)

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

func TestUpdateJob(t *testing.T) {
	ctx := context.Background()
	testID := "654ba8776019d804e35d0d71"
	req := &book.BookRequest{
		Id:   testID,
		From: "customer",
		Desc: "clean",
	}
	res, err := helperService.UpdateJob(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, res.Request.From, req.From)
	assert.Equal(t, res.Request.Desc, req.Desc)
	assert.Equal(t, res.Request.Id, testID)
	assert.NotEmpty(t, res.HelperName)
}

func TestUpdateJobInvalidID(t *testing.T) {
	ctx := context.Background()
	req := &book.BookRequest{
		Id:   "invalid",
		From: "customer",
		Desc: "clean",
	}
	res, err := helperService.UpdateJob(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestSetupFakeData(t *testing.T) {
	helperCollection := dbClient.Database("helper").Collection("helper")
	err := service.SetupFakeData(helperCollection)
	assert.Nil(t, err)

	numDocs, err := helperCollection.CountDocuments(context.Background(), bson.M{})
	assert.Nil(t, err)
	assert.Equal(t, numDocs, int64(3))
}
