package repositories

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/bsgilber/heart-beat/internal/models"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

var ctx = context.TODO()
var db, mock = redismock.NewClientMock()

func TestNewEndpointDb(t *testing.T) {
	type args struct {
		client *redis.Client
	}
	tests := []struct {
		name string
		args args
		want *EndpointDb
	}{
		{
			name: "test happy path creation of endpoint",
			args: args{
				client: db,
			},
			want: &EndpointDb{
				client: db,
				ctx:    ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEndpointDb(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEndpointDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndpointDb_FindIfExists(t *testing.T) {
	t.Run("test happy path where value exists", func(t *testing.T) {
		mock.ExpectSIsMember("endpoints", "some_key").SetVal(true)
		endpointDb := &EndpointDb{
			client: db,
			ctx:    ctx,
		}

		b, e := endpointDb.FindIfExists("some_key")
		assert.Equal(t, true, b)
		assert.Equal(t, nil, e)
	})

	t.Run("test happy path where value does not exist", func(t *testing.T) {
		mock.ExpectSIsMember("endpoints", "some_key").SetVal(false)
		endpointDb := &EndpointDb{
			client: db,
			ctx:    ctx,
		}

		b, e := endpointDb.FindIfExists("some_key")
		assert.Equal(t, false, b)
		assert.Equal(t, nil, e)
	})

	mock.ClearExpect()
}

func TestEndpointDb_FindByName(t *testing.T) {
	t.Run("validate that FindByName returns a populated Endpoint object when it exists in redis", func(t *testing.T) {
		mock.ExpectHGetAll("endpoint-test").SetVal(map[string]string{
			"str1": "test",
			"str2": "https://example.com",
		})

		endpointDb := &EndpointDb{
			client: db,
			ctx:    ctx,
		}

		endpoint, err := endpointDb.FindByName("test")
		assert.Equal(t, "test", endpoint.Name)
		assert.Equal(t, "https://example.com", endpoint.URL)
		assert.Equal(t, nil, err)
	})

	t.Run("validate that FindByName kicks back redis error", func(t *testing.T) {
		mock.ExpectHGetAll("endpoint-test").SetErr(errors.New("error"))

		endpointDb := &EndpointDb{
			client: db,
			ctx:    ctx,
		}

		endpoint, err := endpointDb.FindByName("test")
		assert.Nil(t, endpoint)
		assert.Equal(t, "error", err.Error())
	})

	mock.ClearExpect()
}

func TestEndpointDb_FindAllKeys(t *testing.T) {
	t.Run("validate that FindByName returns a populated Endpoint object when it exists in redis", func(t *testing.T) {
		mock.ExpectSMembers("endpoints").SetVal([]string{"hello", "world"})

		endpointDb := &EndpointDb{
			client: db,
			ctx:    ctx,
		}

		keys := endpointDb.FindAllKeys()
		assert.Equal(t, []string{"hello", "world"}, keys)
	})

	mock.ClearExpect()
}

// TODO: skipping for now, FindAllKeys and FindByName cover core elements
// func TestEndpointDb_FindAll(t *testing.T) {
// }

func TestEndpointDb_Save(t *testing.T) {
	type fields struct {
		client *redis.Client
		ctx    context.Context
	}
	type args struct {
		endpoint models.Endpoint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &EndpointDb{
				client: tt.fields.client,
				ctx:    tt.fields.ctx,
			}
			if err := db.Save(tt.args.endpoint); (err != nil) != tt.wantErr {
				t.Errorf("EndpointDb.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
