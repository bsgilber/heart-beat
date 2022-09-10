package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

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
