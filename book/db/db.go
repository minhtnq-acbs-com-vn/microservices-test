package db

import (
	"context"
	"fmt"
	"github.com/gookit/goutil/errorx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func New(connectionString string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, errorx.Wrap(err, fmt.Sprintf("[DB] Connect Failed %v", err))
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, errorx.Wrap(err, fmt.Sprintf("[DB] Ping Failed %v", err))
	}

	return client, nil
}
