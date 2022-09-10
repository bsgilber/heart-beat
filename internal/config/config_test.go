package config

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		configFilePath string
		env            string
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			name: "test local.yaml",
			args: args{
				configFilePath: "../../config",
				env:            "local",
			},
			want: &Config{
				Server: ServerConfig{
					Port: 8080,
				},
				Logging: LoggerConfig{
					Level: "debug",
				},
				Redis: RedisConnectionConfig{
					Host: "redis://localhost",
					Port: 6379,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.configFilePath, tt.args.env); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
