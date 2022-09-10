package databases

import (
	"context"

	"github.com/bsgilber/heart-beat/internal/models"
	"github.com/go-redis/redis/v8"
)

type EndpointDb struct {
	client *redis.Client
	ctx    context.Context
}

func NewEndpointDb(client *redis.Client) *EndpointDb {
	return &EndpointDb{
		client: client,
		ctx:    context.Background(),
	}
}

func (db *EndpointDb) Save(endpoint models.Endpoint) error {
	_, err := db.client.Pipelined(db.ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(db.ctx, "key", "str1", endpoint.Name)
		rdb.HSet(db.ctx, "key", "str2", endpoint.URL)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (db *EndpointDb) FindByName(name string) (*models.Endpoint, error) {
	var endpoint models.Endpoint

	// Scan all fields into the model.
	err := db.client.HGetAll(db.ctx, "key").Scan(&endpoint)
	if err != nil {
		return nil, err
	}

	return &endpoint, nil
}
