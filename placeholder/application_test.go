package application

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	router := New()
	assert.NotNil(t, router)

	routesInfo := router.srv.Routes()
	assert.NotNil(t, routesInfo)
}

func TestNewBasics(t *testing.T) {
	router := New()

	tests := []struct {
		name     string
		endpoint string
		want     string
	}{
		{
			name:     "test health endpoint",
			endpoint: "/health",
			want:     "Ok",
		},
		{
			name:     "test list endpoint",
			endpoint: "/list",
			want:     "{}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			resp, _ := http.NewRequest(http.MethodGet, tt.endpoint, nil)

			router.srv.ServeHTTP(w, resp)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, tt.want, w.Body.String())

		})
	}
}

func TestInitCache(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "initcache creates empty map",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitCache()
			assert.Equal(t, cache, map[string]string{})
		})
	}
}

func Test_listRegistered(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("List cache objects, empty cache", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		listRegistered(c)
		assert.Equal(t, 200, w.Code)

		b, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, string(b), "{}")
	})
}

func Test_singlePingPrep(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Make sure ping/:name with empty param returns 400", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		singlePingPrep(c)
		assert.Equal(t, 400, w.Code)
	})
	t.Run("Make sure ping/:name with param set returns 200", func(t *testing.T) {
		cache["test"] = "https://example.com"

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// AddParam("id", 1) Result: "/user/1"
		c.AddParam("name", "test")

		singlePingPrep(c)
		assert.Equal(t, 200, w.Code)
	})

}

func Test_allPingPrep(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Make sure pingall with empty cache returns 200", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		allPingPrep(c)
		assert.Equal(t, 200, w.Code)
	})
	t.Run("Make sure pingall with populated cache returns 200", func(t *testing.T) {
		cache["test"] = "https://example.com"

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		allPingPrep(c)
		assert.Equal(t, 200, w.Code)
	})
}

func Test_ping(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Ping endpoint works on happy path", func(t *testing.T) {
		cache["test"] = "https://example.com"

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		ping(c, "test")
		assert.Equal(t, 200, w.Code)

		b, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, string(b), "[test] had status code [200]\n")
	})

	t.Run("Ping endpoint fails when cache key doesnt exist", func(t *testing.T) {
		cache = map[string]string{}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		ping(c, "test")
		assert.Equal(t, 400, w.Code)

		b, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "the key [test] does not exist.\n", string(b))
	})

}

func Test_health(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Health endpoint works on happy path", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		health(c)
		assert.Equal(t, 200, w.Code)

		b, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "Ok", string(b))
	})
}

func Test_register(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Register function returns 400 on empty data payload", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		register(c)
		assert.Equal(t, 400, w.Code)

		b, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "{\"error\":\"invalid request\"}", string(b))
	})

	t.Run("Register function throws 400 on invalid url", func(t *testing.T) {
		cache = map[string]string{}
		router := New()

		body := bytes.NewBuffer([]byte("{\"Name\":\"tester\",\"Url\":\"example.com\"}"))

		req, _ := http.NewRequest("POST", "/register", body)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.srv.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)

		b, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "invalid url, url must contain scheme, host, and/or path", string(b))
	})

	t.Run("Register function adds to cache map on happy path", func(t *testing.T) {
		cache = map[string]string{}
		router := New()

		body := bytes.NewBuffer([]byte("{\"Name\":\"tester\",\"Url\":\"https://example.com\"}"))

		req, _ := http.NewRequest("POST", "/register", body)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.srv.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		b, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "[tester] has been added to the cache.", string(b))
	})
}
