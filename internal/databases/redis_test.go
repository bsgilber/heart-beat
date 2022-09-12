package databases

import (
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestConnectClient(t *testing.T) {
	tests := []struct {
		name string
		want *redis.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnectClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
