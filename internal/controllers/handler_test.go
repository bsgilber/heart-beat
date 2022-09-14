package controllers

import (
	"io/ioutil"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/bsgilber/heart-beat/internal/models"
	"github.com/bsgilber/heart-beat/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

//var ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
var db, mock = redismock.NewClientMock()

func TestNewBaseHandler(t *testing.T) {
	var endpointDb = repositories.NewEndpointDb(db)

	type args struct {
		endpointRepo models.EndpointRepository
	}
	tests := []struct {
		name string
		args args
		want *BaseHandler
	}{
		{
			name: "Happy path where supplied db conn returns a handler with that conn.",
			args: args{
				endpointRepo: endpointDb,
			},
			want: &BaseHandler{
				endpointRepo: endpointDb,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBaseHandler(tt.args.endpointRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaseHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseHandler_Ping(t *testing.T) {
	var endpointDb = repositories.NewEndpointDb(db)

	t.Run("test error path where name/key does not exist", func(t *testing.T) {
		baseHandler := &BaseHandler{
			endpointRepo: endpointDb,
		}

		mock.ExpectSIsMember("endpoints", "test").SetVal(false)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		baseHandler.Ping(c, "test")

		b, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "the key [test] does not exist.\n", string(b))
	})

	mock.ClearExpect()

	t.Run("test happy path where name/key does exist", func(t *testing.T) {
		baseHandler := &BaseHandler{
			endpointRepo: endpointDb,
		}

		mock.ExpectSIsMember("endpoints", "test").SetVal(true)
		mock.ExpectHGetAll("endpoint-test").SetVal(map[string]string{
			"str1": "test",
			"str2": "https://example.com",
		})

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		baseHandler.Ping(c, "test")

		b, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "[test] had status code [200]\n", string(b))
	})

	mock.ClearExpect()

	t.Run("test path where name/key does exist but a bad url is provided", func(t *testing.T) {
		baseHandler := &BaseHandler{
			endpointRepo: endpointDb,
		}

		mock.ExpectSIsMember("endpoints", "test").SetVal(true)
		mock.ExpectHGetAll("endpoint-test").SetVal(map[string]string{
			"str1": "test",
			"str2": "example.com",
		})

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		baseHandler.Ping(c, "test")

		b, _ := ioutil.ReadAll(w.Body)
		assert.Contains(t, string(b), "unsupported protocol")
	})

	mock.ClearExpect()
}

func TestBaseHandler_Register(t *testing.T) {
	var endpointDb = repositories.NewEndpointDb(db)

	t.Run("test error when no data request object is passed in", func(t *testing.T) {
		baseHandler := &BaseHandler{
			endpointRepo: endpointDb,
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		baseHandler.Register(c)

		b, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "{\"error\":\"invalid request\"}", string(b))
	})

	//t.Run("test error path where name/key does not exist", func(t *testing.T) {
	//	baseHandler := &BaseHandler{
	//		endpointRepo: endpointDb,
	//	}

	//	mock.ExpectSIsMember("endpoints", "test").SetVal(true)
	//	w := httptest.NewRecorder()
	//	c, _ := gin.CreateTestContext(w)

	//	baseHandler.Register(c)

	//	b, _ := ioutil.ReadAll(w.Body)
	//	assert.Equal(t, "name must be unique; [test] provided, exist check returned [false].", string(b))
	//})

	//mock.ClearExpect()
}

func TestBaseHandler_SinglePingPrep(t *testing.T) {
	type fields struct {
		endpointRepo models.EndpointRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &BaseHandler{
				endpointRepo: tt.fields.endpointRepo,
			}
			h.SinglePingPrep(tt.args.c)
		})
	}
}

func TestBaseHandler_AllPingPrep(t *testing.T) {
	type fields struct {
		endpointRepo models.EndpointRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &BaseHandler{
				endpointRepo: tt.fields.endpointRepo,
			}
			h.AllPingPrep(tt.args.c)
		})
	}
}

func TestBaseHandler_ListRegistered(t *testing.T) {
	type fields struct {
		endpointRepo models.EndpointRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &BaseHandler{
				endpointRepo: tt.fields.endpointRepo,
			}
			h.ListRegistered(tt.args.c)
		})
	}
}
