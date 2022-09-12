package repositories

import (
	"context"
	"fmt"

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

func (db *EndpointDb) FindIfExists(name string) (bool, error) {
	is, err := db.client.SIsMember(db.ctx, "endpoints", name).Result()
	if err != nil {
		return false, err
	}

	return is, nil
}

func (db *EndpointDb) FindByName(name string) (*models.Endpoint, error) {
	var endpoint models.Endpoint

	// Scan all fields into the model.
	err := db.client.HGetAll(db.ctx, fmt.Sprintf("endpoint-%s", name)).Scan(&endpoint)
	if err != nil {
		return nil, err
	}

	return &endpoint, nil
}

func (db *EndpointDb) FindAll() []*models.Endpoint {
	var endpoints []*models.Endpoint

	members := db.FindAllKeys()

	for _, member := range members {
		val, _ := db.FindByName(member)
		endpoints = append(endpoints, val)
	}

	return endpoints
}

func (db *EndpointDb) FindAllKeys() []string {
	return db.client.SMembers(db.ctx, "endpoints").Val()
}

func (db *EndpointDb) Save(endpoint models.Endpoint) error {
	_, err := db.client.Pipelined(db.ctx, func(rdb redis.Pipeliner) error {
		// create a hash object for the endpoint name
		rdb.HSet(db.ctx, fmt.Sprintf("endpoint-%s", endpoint.Name), "str1", endpoint.Name)
		rdb.HSet(db.ctx, fmt.Sprintf("endpoint-%s", endpoint.Name), "str2", endpoint.URL)

		// add the hash key to a set so we can get a list of all keys easily
		rdb.SAdd(db.ctx, "endpoints", "endpoint-%s", endpoint.Name)

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
