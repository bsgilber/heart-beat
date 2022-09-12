package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	router := New()
	assert.NotNil(t, router)

	routesInfo := router.srv.Routes()
	assert.NotNil(t, routesInfo)
}

func TestInitEnv(t *testing.T) {
	InitEnv()

	assert.Equal(t, environment, "local")
}
