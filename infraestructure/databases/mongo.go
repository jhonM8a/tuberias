package databases

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnector struct {
	client *mongo.Client
}

func (m *MongoDBConnector) Connect(dsn string) (interface{}, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	m.client = client
	return client, nil
}

func (m *MongoDBConnector) Insert(collection string, document interface{}) error {
	if m.client == nil {
		return mongo.ErrClientDisconnected
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := m.client.Database("proyecto").Collection(collection)
	_, err := coll.InsertOne(ctx, document)
	return err
}
