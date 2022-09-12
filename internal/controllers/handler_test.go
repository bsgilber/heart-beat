package controllers

import (
	"reflect"
	"testing"

	"github.com/bsgilber/heart-beat/internal/models"
	"github.com/gin-gonic/gin"
)

func TestNewBaseHandler(t *testing.T) {
	type args struct {
		endpointRepo models.EndpointRepository
	}
	tests := []struct {
		name string
		args args
		want *BaseHandler
	}{
		// TODO: Add test cases.
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
	type fields struct {
		endpointRepo models.EndpointRepository
	}
	type args struct {
		c    *gin.Context
		name string
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
			h.Ping(tt.args.c, tt.args.name)
		})
	}
}

func TestBaseHandler_Register(t *testing.T) {
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
			h.Register(tt.args.c)
		})
	}
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
