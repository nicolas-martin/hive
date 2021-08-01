package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ConnectionString string `envconfig:"CONNECTION_STRING"`
}

// Load parses the env varaibles into a config struct
func Load() (*Config, error) {
	cfg := Config{}

	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error processing env config: %w", err)
	}
	return &cfg, nil
}
